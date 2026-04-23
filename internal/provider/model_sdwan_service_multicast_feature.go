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
type ServiceMulticast struct {
	Id                                           types.String                         `tfsdk:"id"`
	Version                                      types.Int64                          `tfsdk:"version"`
	Name                                         types.String                         `tfsdk:"name"`
	Description                                  types.String                         `tfsdk:"description"`
	FeatureProfileId                             types.String                         `tfsdk:"feature_profile_id"`
	SptOnly                                      types.Bool                           `tfsdk:"spt_only"`
	SptOnlyVariable                              types.String                         `tfsdk:"spt_only_variable"`
	LocalReplicator                              types.Bool                           `tfsdk:"local_replicator"`
	LocalReplicatorVariable                      types.String                         `tfsdk:"local_replicator_variable"`
	LocalReplicatorThreshold                     types.Int64                          `tfsdk:"local_replicator_threshold"`
	LocalReplicatorThresholdVariable             types.String                         `tfsdk:"local_replicator_threshold_variable"`
	IgmpInterfaces                               []ServiceMulticastIgmpInterfaces     `tfsdk:"igmp_interfaces"`
	PimSourceSpecificMulticastEnable             types.Bool                           `tfsdk:"pim_source_specific_multicast_enable"`
	PimSourceSpecificMulticastAccessList         types.String                         `tfsdk:"pim_source_specific_multicast_access_list"`
	PimSourceSpecificMulticastAccessListVariable types.String                         `tfsdk:"pim_source_specific_multicast_access_list_variable"`
	PimSptThreshold                              types.String                         `tfsdk:"pim_spt_threshold"`
	PimSptThresholdVariable                      types.String                         `tfsdk:"pim_spt_threshold_variable"`
	PimInterfaces                                []ServiceMulticastPimInterfaces      `tfsdk:"pim_interfaces"`
	StaticRpAddresses                            []ServiceMulticastStaticRpAddresses  `tfsdk:"static_rp_addresses"`
	EnableAutoRp                                 types.Bool                           `tfsdk:"enable_auto_rp"`
	EnableAutoRpVariable                         types.String                         `tfsdk:"enable_auto_rp_variable"`
	AutoRpAnnounces                              []ServiceMulticastAutoRpAnnounces    `tfsdk:"auto_rp_announces"`
	AutoRpDiscoveries                            []ServiceMulticastAutoRpDiscoveries  `tfsdk:"auto_rp_discoveries"`
	PimBsrRpCandidates                           []ServiceMulticastPimBsrRpCandidates `tfsdk:"pim_bsr_rp_candidates"`
	PimBsrCandidates                             []ServiceMulticastPimBsrCandidates   `tfsdk:"pim_bsr_candidates"`
	MsdpGroups                                   []ServiceMulticastMsdpGroups         `tfsdk:"msdp_groups"`
	MsdpOriginatorId                             types.String                         `tfsdk:"msdp_originator_id"`
	MsdpOriginatorIdVariable                     types.String                         `tfsdk:"msdp_originator_id_variable"`
	MsdpConnectionRetryInterval                  types.Int64                          `tfsdk:"msdp_connection_retry_interval"`
	MsdpConnectionRetryIntervalVariable          types.String                         `tfsdk:"msdp_connection_retry_interval_variable"`
}

type ServiceMulticastIgmpInterfaces struct {
	InterfaceName         types.String                               `tfsdk:"interface_name"`
	InterfaceNameVariable types.String                               `tfsdk:"interface_name_variable"`
	Version               types.Int64                                `tfsdk:"version"`
	JoinGroups            []ServiceMulticastIgmpInterfacesJoinGroups `tfsdk:"join_groups"`
}

type ServiceMulticastPimInterfaces struct {
	InterfaceName             types.String `tfsdk:"interface_name"`
	InterfaceNameVariable     types.String `tfsdk:"interface_name_variable"`
	QueryInterval             types.Int64  `tfsdk:"query_interval"`
	QueryIntervalVariable     types.String `tfsdk:"query_interval_variable"`
	JoinPruneInterval         types.Int64  `tfsdk:"join_prune_interval"`
	JoinPruneIntervalVariable types.String `tfsdk:"join_prune_interval_variable"`
}

type ServiceMulticastStaticRpAddresses struct {
	IpAddress          types.String `tfsdk:"ip_address"`
	IpAddressVariable  types.String `tfsdk:"ip_address_variable"`
	AccessList         types.String `tfsdk:"access_list"`
	AccessListVariable types.String `tfsdk:"access_list_variable"`
	Override           types.Bool   `tfsdk:"override"`
	OverrideVariable   types.String `tfsdk:"override_variable"`
}

type ServiceMulticastAutoRpAnnounces struct {
	InterfaceName         types.String `tfsdk:"interface_name"`
	InterfaceNameVariable types.String `tfsdk:"interface_name_variable"`
	Scope                 types.Int64  `tfsdk:"scope"`
	ScopeVariable         types.String `tfsdk:"scope_variable"`
}

type ServiceMulticastAutoRpDiscoveries struct {
	InterfaceName         types.String `tfsdk:"interface_name"`
	InterfaceNameVariable types.String `tfsdk:"interface_name_variable"`
	Scope                 types.Int64  `tfsdk:"scope"`
	ScopeVariable         types.String `tfsdk:"scope_variable"`
}

type ServiceMulticastPimBsrRpCandidates struct {
	InterfaceName         types.String `tfsdk:"interface_name"`
	InterfaceNameVariable types.String `tfsdk:"interface_name_variable"`
	AccessListId          types.String `tfsdk:"access_list_id"`
	AccessListIdVariable  types.String `tfsdk:"access_list_id_variable"`
	Interval              types.Int64  `tfsdk:"interval"`
	IntervalVariable      types.String `tfsdk:"interval_variable"`
	Priority              types.Int64  `tfsdk:"priority"`
	PriorityVariable      types.String `tfsdk:"priority_variable"`
}

type ServiceMulticastPimBsrCandidates struct {
	InterfaceName                     types.String `tfsdk:"interface_name"`
	InterfaceNameVariable             types.String `tfsdk:"interface_name_variable"`
	HashMaskLength                    types.Int64  `tfsdk:"hash_mask_length"`
	HashMaskLengthVariable            types.String `tfsdk:"hash_mask_length_variable"`
	Priority                          types.Int64  `tfsdk:"priority"`
	PriorityVariable                  types.String `tfsdk:"priority_variable"`
	AcceptCandidateAccessList         types.String `tfsdk:"accept_candidate_access_list"`
	AcceptCandidateAccessListVariable types.String `tfsdk:"accept_candidate_access_list_variable"`
}

type ServiceMulticastMsdpGroups struct {
	MeshGroupName         types.String                      `tfsdk:"mesh_group_name"`
	MeshGroupNameVariable types.String                      `tfsdk:"mesh_group_name_variable"`
	Peers                 []ServiceMulticastMsdpGroupsPeers `tfsdk:"peers"`
}

type ServiceMulticastIgmpInterfacesJoinGroups struct {
	GroupAddress          types.String `tfsdk:"group_address"`
	GroupAddressVariable  types.String `tfsdk:"group_address_variable"`
	SourceAddress         types.String `tfsdk:"source_address"`
	SourceAddressVariable types.String `tfsdk:"source_address_variable"`
}

type ServiceMulticastMsdpGroupsPeers struct {
	PeerIp                             types.String `tfsdk:"peer_ip"`
	PeerIpVariable                     types.String `tfsdk:"peer_ip_variable"`
	ConnectionSourceInterface          types.String `tfsdk:"connection_source_interface"`
	ConnectionSourceInterfaceVariable  types.String `tfsdk:"connection_source_interface_variable"`
	RemoteAs                           types.Int64  `tfsdk:"remote_as"`
	RemoteAsVariable                   types.String `tfsdk:"remote_as_variable"`
	PeerAuthenticationPassword         types.String `tfsdk:"peer_authentication_password"`
	PeerAuthenticationPasswordVariable types.String `tfsdk:"peer_authentication_password_variable"`
	KeepaliveInterval                  types.Int64  `tfsdk:"keepalive_interval"`
	KeepaliveIntervalVariable          types.String `tfsdk:"keepalive_interval_variable"`
	KeepaliveHoldTime                  types.Int64  `tfsdk:"keepalive_hold_time"`
	KeepaliveHoldTimeVariable          types.String `tfsdk:"keepalive_hold_time_variable"`
	SaLimit                            types.Int64  `tfsdk:"sa_limit"`
	SaLimitVariable                    types.String `tfsdk:"sa_limit_variable"`
	DefaultPeer                        types.Bool   `tfsdk:"default_peer"`
	PrefixListId                       types.String `tfsdk:"prefix_list_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceMulticast) getModel() string {
	return "service_multicast"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceMulticast) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/routing/multicast", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceMulticast) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.SptOnlyVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.sptOnly.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.sptOnly.value", data.SptOnlyVariable.ValueString())
		}
	} else if data.SptOnly.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.sptOnly.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.sptOnly.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.sptOnly.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.sptOnly.value", data.SptOnly.ValueBool())
		}
	}

	if !data.LocalReplicatorVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.localConfig.local.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.localConfig.local.value", data.LocalReplicatorVariable.ValueString())
		}
	} else if data.LocalReplicator.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.localConfig.local.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.localConfig.local.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.localConfig.local.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.localConfig.local.value", data.LocalReplicator.ValueBool())
		}
	}

	if !data.LocalReplicatorThresholdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.localConfig.threshold.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.localConfig.threshold.value", data.LocalReplicatorThresholdVariable.ValueString())
		}
	} else if data.LocalReplicatorThreshold.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.localConfig.threshold.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.localConfig.threshold.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.localConfig.threshold.value", data.LocalReplicatorThreshold.ValueInt64())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"igmp.interface", []interface{}{})
		for _, item := range data.IgmpInterfaces {
			itemBody := ""

			if !item.InterfaceNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceNameVariable.ValueString())
				}
			} else if !item.InterfaceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceName.ValueString())
				}
			}
			if item.Version.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "version.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "version.value", 2)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "version.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "version.value", item.Version.ValueInt64())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "joinGroup", []interface{}{})
				for _, childItem := range item.JoinGroups {
					itemChildBody := ""

					if !childItem.GroupAddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "groupAddress.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "groupAddress.value", childItem.GroupAddressVariable.ValueString())
						}
					} else if !childItem.GroupAddress.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "groupAddress.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "groupAddress.value", childItem.GroupAddress.ValueString())
						}
					}

					if !childItem.SourceAddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceAddress.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceAddress.value", childItem.SourceAddressVariable.ValueString())
						}
					} else if childItem.SourceAddress.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceAddress.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceAddress.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceAddress.value", childItem.SourceAddress.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "joinGroup.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"igmp.interface.-1", itemBody)
		}
	}
	if !data.PimSourceSpecificMulticastEnable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"pim.ssm.ssmRangeConfig.enableSSMFlag.optionType", "global")
			body, _ = sjson.Set(body, path+"pim.ssm.ssmRangeConfig.enableSSMFlag.value", data.PimSourceSpecificMulticastEnable.ValueBool())
		}
	}

	if !data.PimSourceSpecificMulticastAccessListVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"pim.ssm.ssmRangeConfig.range.optionType", "variable")
			body, _ = sjson.Set(body, path+"pim.ssm.ssmRangeConfig.range.value", data.PimSourceSpecificMulticastAccessListVariable.ValueString())
		}
	} else if !data.PimSourceSpecificMulticastAccessList.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"pim.ssm.ssmRangeConfig.range.optionType", "global")
			body, _ = sjson.Set(body, path+"pim.ssm.ssmRangeConfig.range.value", data.PimSourceSpecificMulticastAccessList.ValueString())
		}
	}

	if !data.PimSptThresholdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"pim.ssm.sptThreshold.optionType", "variable")
			body, _ = sjson.Set(body, path+"pim.ssm.sptThreshold.value", data.PimSptThresholdVariable.ValueString())
		}
	} else if data.PimSptThreshold.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"pim.ssm.sptThreshold.optionType", "default")
			body, _ = sjson.Set(body, path+"pim.ssm.sptThreshold.value", "0")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"pim.ssm.sptThreshold.optionType", "global")
			body, _ = sjson.Set(body, path+"pim.ssm.sptThreshold.value", data.PimSptThreshold.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"pim.interface", []interface{}{})
		for _, item := range data.PimInterfaces {
			itemBody := ""

			if !item.InterfaceNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceNameVariable.ValueString())
				}
			} else if !item.InterfaceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceName.ValueString())
				}
			}

			if !item.QueryIntervalVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "queryInterval.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "queryInterval.value", item.QueryIntervalVariable.ValueString())
				}
			} else if item.QueryInterval.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "queryInterval.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "queryInterval.value", 30)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "queryInterval.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "queryInterval.value", item.QueryInterval.ValueInt64())
				}
			}

			if !item.JoinPruneIntervalVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "joinPruneInterval.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "joinPruneInterval.value", item.JoinPruneIntervalVariable.ValueString())
				}
			} else if item.JoinPruneInterval.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "joinPruneInterval.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "joinPruneInterval.value", 60)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "joinPruneInterval.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "joinPruneInterval.value", item.JoinPruneInterval.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"pim.interface.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"pim.rpAddr", []interface{}{})
		for _, item := range data.StaticRpAddresses {
			itemBody := ""

			if !item.IpAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.IpAddressVariable.ValueString())
				}
			} else if !item.IpAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.IpAddress.ValueString())
				}
			}

			if !item.AccessListVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "accessList.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "accessList.value", item.AccessListVariable.ValueString())
				}
			} else if !item.AccessList.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "accessList.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "accessList.value", item.AccessList.ValueString())
				}
			}

			if !item.OverrideVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "override.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "override.value", item.OverrideVariable.ValueString())
				}
			} else if item.Override.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "override.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "override.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "override.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "override.value", item.Override.ValueBool())
				}
			}
			body, _ = sjson.SetRaw(body, path+"pim.rpAddr.-1", itemBody)
		}
	}

	if !data.EnableAutoRpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"pim.autoRp.enableAutoRPFlag.optionType", "variable")
			body, _ = sjson.Set(body, path+"pim.autoRp.enableAutoRPFlag.value", data.EnableAutoRpVariable.ValueString())
		}
	} else if data.EnableAutoRp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"pim.autoRp.enableAutoRPFlag.optionType", "default")
			body, _ = sjson.Set(body, path+"pim.autoRp.enableAutoRPFlag.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"pim.autoRp.enableAutoRPFlag.optionType", "global")
			body, _ = sjson.Set(body, path+"pim.autoRp.enableAutoRPFlag.value", data.EnableAutoRp.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"pim.autoRp.sendRpAnnounceList", []interface{}{})
		for _, item := range data.AutoRpAnnounces {
			itemBody := ""

			if !item.InterfaceNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceNameVariable.ValueString())
				}
			} else if !item.InterfaceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceName.ValueString())
				}
			}

			if !item.ScopeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "scope.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "scope.value", item.ScopeVariable.ValueString())
				}
			} else if !item.Scope.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "scope.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "scope.value", item.Scope.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"pim.autoRp.sendRpAnnounceList.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"pim.autoRp.sendRpDiscovery", []interface{}{})
		for _, item := range data.AutoRpDiscoveries {
			itemBody := ""

			if !item.InterfaceNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceNameVariable.ValueString())
				}
			} else if !item.InterfaceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceName.ValueString())
				}
			}

			if !item.ScopeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "scope.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "scope.value", item.ScopeVariable.ValueString())
				}
			} else if !item.Scope.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "scope.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "scope.value", item.Scope.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"pim.autoRp.sendRpDiscovery.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"pim.pimBsr.rpCandidate", []interface{}{})
		for _, item := range data.PimBsrRpCandidates {
			itemBody := ""

			if !item.InterfaceNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceNameVariable.ValueString())
				}
			} else if !item.InterfaceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceName.ValueString())
				}
			}

			if !item.AccessListIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupList.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "groupList.value", item.AccessListIdVariable.ValueString())
				}
			} else if item.AccessListId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupList.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupList.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "groupList.value", item.AccessListId.ValueString())
				}
			}

			if !item.IntervalVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interval.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "interval.value", item.IntervalVariable.ValueString())
				}
			} else if item.Interval.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interval.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interval.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "interval.value", item.Interval.ValueInt64())
				}
			}

			if !item.PriorityVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.PriorityVariable.ValueString())
				}
			} else if item.Priority.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.Priority.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"pim.pimBsr.rpCandidate.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"pim.pimBsr.bsrCandidate", []interface{}{})
		for _, item := range data.PimBsrCandidates {
			itemBody := ""

			if !item.InterfaceNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceNameVariable.ValueString())
				}
			} else if !item.InterfaceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interfaceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "interfaceName.value", item.InterfaceName.ValueString())
				}
			}

			if !item.HashMaskLengthVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "mask.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "mask.value", item.HashMaskLengthVariable.ValueString())
				}
			} else if item.HashMaskLength.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "mask.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "mask.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "mask.value", item.HashMaskLength.ValueInt64())
				}
			}

			if !item.PriorityVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.PriorityVariable.ValueString())
				}
			} else if item.Priority.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.Priority.ValueInt64())
				}
			}

			if !item.AcceptCandidateAccessListVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "acceptRpCandidate.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "acceptRpCandidate.value", item.AcceptCandidateAccessListVariable.ValueString())
				}
			} else if item.AcceptCandidateAccessList.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "acceptRpCandidate.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "acceptRpCandidate.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "acceptRpCandidate.value", item.AcceptCandidateAccessList.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"pim.pimBsr.bsrCandidate.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"msdp.msdpList", []interface{}{})
		for _, item := range data.MsdpGroups {
			itemBody := ""

			if !item.MeshGroupNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "meshGroup.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "meshGroup.value", item.MeshGroupNameVariable.ValueString())
				}
			} else if item.MeshGroupName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "meshGroup.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "meshGroup.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "meshGroup.value", item.MeshGroupName.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "peer", []interface{}{})
				for _, childItem := range item.Peers {
					itemChildBody := ""

					if !childItem.PeerIpVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "peerIp.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "peerIp.value", childItem.PeerIpVariable.ValueString())
						}
					} else if !childItem.PeerIp.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "peerIp.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "peerIp.value", childItem.PeerIp.ValueString())
						}
					}

					if !childItem.ConnectionSourceInterfaceVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "connectSourceIntf.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "connectSourceIntf.value", childItem.ConnectionSourceInterfaceVariable.ValueString())
						}
					} else if childItem.ConnectionSourceInterface.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "connectSourceIntf.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "connectSourceIntf.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "connectSourceIntf.value", childItem.ConnectionSourceInterface.ValueString())
						}
					}

					if !childItem.RemoteAsVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "remoteAs.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "remoteAs.value", childItem.RemoteAsVariable.ValueString())
						}
					} else if childItem.RemoteAs.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "remoteAs.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "remoteAs.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "remoteAs.value", childItem.RemoteAs.ValueInt64())
						}
					}

					if !childItem.PeerAuthenticationPasswordVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "password.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "password.value", childItem.PeerAuthenticationPasswordVariable.ValueString())
						}
					} else if childItem.PeerAuthenticationPassword.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "password.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "password.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "password.value", childItem.PeerAuthenticationPassword.ValueString())
						}
					}

					if !childItem.KeepaliveIntervalVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keepaliveInterval.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "keepaliveInterval.value", childItem.KeepaliveIntervalVariable.ValueString())
						}
					} else if childItem.KeepaliveInterval.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keepaliveInterval.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keepaliveInterval.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "keepaliveInterval.value", childItem.KeepaliveInterval.ValueInt64())
						}
					}

					if !childItem.KeepaliveHoldTimeVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keepaliveHoldTime.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "keepaliveHoldTime.value", childItem.KeepaliveHoldTimeVariable.ValueString())
						}
					} else if childItem.KeepaliveHoldTime.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keepaliveHoldTime.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keepaliveHoldTime.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "keepaliveHoldTime.value", childItem.KeepaliveHoldTime.ValueInt64())
						}
					}

					if !childItem.SaLimitVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "saLimit.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "saLimit.value", childItem.SaLimitVariable.ValueString())
						}
					} else if childItem.SaLimit.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "saLimit.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "saLimit.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "saLimit.value", childItem.SaLimit.ValueInt64())
						}
					}
					if !childItem.DefaultPeer.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "default.defaultPeer.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "default.defaultPeer.value", childItem.DefaultPeer.ValueBool())
						}
					}
					if !childItem.PrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "default.prefixList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "default.prefixList.refId.value", childItem.PrefixListId.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "peer.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"msdp.msdpList.-1", itemBody)
		}
	}

	if !data.MsdpOriginatorIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"msdp.originatorId.optionType", "variable")
			body, _ = sjson.Set(body, path+"msdp.originatorId.value", data.MsdpOriginatorIdVariable.ValueString())
		}
	} else if data.MsdpOriginatorId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"msdp.originatorId.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"msdp.originatorId.optionType", "global")
			body, _ = sjson.Set(body, path+"msdp.originatorId.value", data.MsdpOriginatorId.ValueString())
		}
	}

	if !data.MsdpConnectionRetryIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"msdp.refreshTimer.optionType", "variable")
			body, _ = sjson.Set(body, path+"msdp.refreshTimer.value", data.MsdpConnectionRetryIntervalVariable.ValueString())
		}
	} else if data.MsdpConnectionRetryInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"msdp.refreshTimer.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"msdp.refreshTimer.optionType", "global")
			body, _ = sjson.Set(body, path+"msdp.refreshTimer.value", data.MsdpConnectionRetryInterval.ValueInt64())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceMulticast) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.SptOnly = types.BoolNull()
	data.SptOnlyVariable = types.StringNull()
	if t := res.Get(path + "basic.sptOnly.optionType"); t.Exists() {
		va := res.Get(path + "basic.sptOnly.value")
		if t.String() == "variable" {
			data.SptOnlyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SptOnly = types.BoolValue(va.Bool())
		}
	}
	data.LocalReplicator = types.BoolNull()
	data.LocalReplicatorVariable = types.StringNull()
	if t := res.Get(path + "basic.localConfig.local.optionType"); t.Exists() {
		va := res.Get(path + "basic.localConfig.local.value")
		if t.String() == "variable" {
			data.LocalReplicatorVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LocalReplicator = types.BoolValue(va.Bool())
		}
	}
	data.LocalReplicatorThreshold = types.Int64Null()
	data.LocalReplicatorThresholdVariable = types.StringNull()
	if t := res.Get(path + "basic.localConfig.threshold.optionType"); t.Exists() {
		va := res.Get(path + "basic.localConfig.threshold.value")
		if t.String() == "variable" {
			data.LocalReplicatorThresholdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LocalReplicatorThreshold = types.Int64Value(va.Int())
		}
	}
	oldIgmpInterfaces := data.IgmpInterfaces
	if value := res.Get(path + "igmp.interface"); value.Exists() && len(value.Array()) > 0 {
		data.IgmpInterfaces = make([]ServiceMulticastIgmpInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceMulticastIgmpInterfaces{}
			item.InterfaceName = types.StringNull()
			item.InterfaceNameVariable = types.StringNull()
			if t := v.Get("interfaceName.optionType"); t.Exists() {
				va := v.Get("interfaceName.value")
				if t.String() == "variable" {
					item.InterfaceNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.InterfaceName = types.StringValue(va.String())
				}
			}
			item.Version = types.Int64Null()

			if t := v.Get("version.optionType"); t.Exists() {
				va := v.Get("version.value")
				if t.String() == "global" {
					item.Version = types.Int64Value(va.Int())
				}
			}
			if cValue := v.Get("joinGroup"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.JoinGroups = make([]ServiceMulticastIgmpInterfacesJoinGroups, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceMulticastIgmpInterfacesJoinGroups{}
					cItem.GroupAddress = types.StringNull()
					cItem.GroupAddressVariable = types.StringNull()
					if t := cv.Get("groupAddress.optionType"); t.Exists() {
						va := cv.Get("groupAddress.value")
						if t.String() == "variable" {
							cItem.GroupAddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.GroupAddress = types.StringValue(va.String())
						}
					}
					cItem.SourceAddress = types.StringNull()
					cItem.SourceAddressVariable = types.StringNull()
					if t := cv.Get("sourceAddress.optionType"); t.Exists() {
						va := cv.Get("sourceAddress.value")
						if t.String() == "variable" {
							cItem.SourceAddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.SourceAddress = types.StringValue(va.String())
						}
					}
					item.JoinGroups = append(item.JoinGroups, cItem)
					return true
				})
			}
			data.IgmpInterfaces = append(data.IgmpInterfaces, item)
			return true
		})
	} else {
		data.IgmpInterfaces = nil
	}
	if !fullRead {
		resultIgmpInterfaces := make([]ServiceMulticastIgmpInterfaces, 0, len(data.IgmpInterfaces))
		matchedIgmpInterfaces := make([]bool, len(data.IgmpInterfaces))
		for _, oldItem := range oldIgmpInterfaces {
			for ni := range data.IgmpInterfaces {
				if matchedIgmpInterfaces[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.InterfaceNameVariable.ValueString() != "" || data.IgmpInterfaces[ni].InterfaceNameVariable.ValueString() != "") {
					if oldItem.InterfaceNameVariable.ValueString() != data.IgmpInterfaces[ni].InterfaceNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.InterfaceName.ValueString() != data.IgmpInterfaces[ni].InterfaceName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedIgmpInterfaces[ni] = true
					{
						resultC := make([]ServiceMulticastIgmpInterfacesJoinGroups, 0, len(data.IgmpInterfaces[ni].JoinGroups))
						matchedC := make([]bool, len(data.IgmpInterfaces[ni].JoinGroups))
						for _, oldCItem := range oldItem.JoinGroups {
							for nci := range data.IgmpInterfaces[ni].JoinGroups {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.GroupAddressVariable.ValueString() != "" || data.IgmpInterfaces[ni].JoinGroups[nci].GroupAddressVariable.ValueString() != "") {
									if oldCItem.GroupAddressVariable.ValueString() != data.IgmpInterfaces[ni].JoinGroups[nci].GroupAddressVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.GroupAddress.ValueString() != data.IgmpInterfaces[ni].JoinGroups[nci].GroupAddress.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC && (oldCItem.SourceAddressVariable.ValueString() != "" || data.IgmpInterfaces[ni].JoinGroups[nci].SourceAddressVariable.ValueString() != "") {
									if oldCItem.SourceAddressVariable.ValueString() != data.IgmpInterfaces[ni].JoinGroups[nci].SourceAddressVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.SourceAddress.ValueString() != data.IgmpInterfaces[ni].JoinGroups[nci].SourceAddress.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.IgmpInterfaces[ni].JoinGroups[nci])
									break
								}
							}
						}
						for nci := range data.IgmpInterfaces[ni].JoinGroups {
							if !matchedC[nci] {
								resultC = append(resultC, data.IgmpInterfaces[ni].JoinGroups[nci])
							}
						}
						data.IgmpInterfaces[ni].JoinGroups = resultC
					}
					resultIgmpInterfaces = append(resultIgmpInterfaces, data.IgmpInterfaces[ni])
					break
				}
			}
		}
		for ni := range data.IgmpInterfaces {
			if !matchedIgmpInterfaces[ni] {
				resultIgmpInterfaces = append(resultIgmpInterfaces, data.IgmpInterfaces[ni])
			}
		}
		data.IgmpInterfaces = resultIgmpInterfaces
	}
	data.PimSourceSpecificMulticastEnable = types.BoolNull()

	if t := res.Get(path + "pim.ssm.ssmRangeConfig.enableSSMFlag.optionType"); t.Exists() {
		va := res.Get(path + "pim.ssm.ssmRangeConfig.enableSSMFlag.value")
		if t.String() == "global" {
			data.PimSourceSpecificMulticastEnable = types.BoolValue(va.Bool())
		}
	}
	data.PimSourceSpecificMulticastAccessList = types.StringNull()
	data.PimSourceSpecificMulticastAccessListVariable = types.StringNull()
	if t := res.Get(path + "pim.ssm.ssmRangeConfig.range.optionType"); t.Exists() {
		va := res.Get(path + "pim.ssm.ssmRangeConfig.range.value")
		if t.String() == "variable" {
			data.PimSourceSpecificMulticastAccessListVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PimSourceSpecificMulticastAccessList = types.StringValue(va.String())
		}
	}
	data.PimSptThreshold = types.StringNull()
	data.PimSptThresholdVariable = types.StringNull()
	if t := res.Get(path + "pim.ssm.sptThreshold.optionType"); t.Exists() {
		va := res.Get(path + "pim.ssm.sptThreshold.value")
		if t.String() == "variable" {
			data.PimSptThresholdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PimSptThreshold = types.StringValue(va.String())
		}
	}
	oldPimInterfaces := data.PimInterfaces
	if value := res.Get(path + "pim.interface"); value.Exists() && len(value.Array()) > 0 {
		data.PimInterfaces = make([]ServiceMulticastPimInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceMulticastPimInterfaces{}
			item.InterfaceName = types.StringNull()
			item.InterfaceNameVariable = types.StringNull()
			if t := v.Get("interfaceName.optionType"); t.Exists() {
				va := v.Get("interfaceName.value")
				if t.String() == "variable" {
					item.InterfaceNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.InterfaceName = types.StringValue(va.String())
				}
			}
			item.QueryInterval = types.Int64Null()
			item.QueryIntervalVariable = types.StringNull()
			if t := v.Get("queryInterval.optionType"); t.Exists() {
				va := v.Get("queryInterval.value")
				if t.String() == "variable" {
					item.QueryIntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.QueryInterval = types.Int64Value(va.Int())
				}
			}
			item.JoinPruneInterval = types.Int64Null()
			item.JoinPruneIntervalVariable = types.StringNull()
			if t := v.Get("joinPruneInterval.optionType"); t.Exists() {
				va := v.Get("joinPruneInterval.value")
				if t.String() == "variable" {
					item.JoinPruneIntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.JoinPruneInterval = types.Int64Value(va.Int())
				}
			}
			data.PimInterfaces = append(data.PimInterfaces, item)
			return true
		})
	} else {
		data.PimInterfaces = nil
	}
	if !fullRead {
		resultPimInterfaces := make([]ServiceMulticastPimInterfaces, 0, len(data.PimInterfaces))
		matchedPimInterfaces := make([]bool, len(data.PimInterfaces))
		for _, oldItem := range oldPimInterfaces {
			for ni := range data.PimInterfaces {
				if matchedPimInterfaces[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.InterfaceNameVariable.ValueString() != "" || data.PimInterfaces[ni].InterfaceNameVariable.ValueString() != "") {
					if oldItem.InterfaceNameVariable.ValueString() != data.PimInterfaces[ni].InterfaceNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.InterfaceName.ValueString() != data.PimInterfaces[ni].InterfaceName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedPimInterfaces[ni] = true
					resultPimInterfaces = append(resultPimInterfaces, data.PimInterfaces[ni])
					break
				}
			}
		}
		for ni := range data.PimInterfaces {
			if !matchedPimInterfaces[ni] {
				resultPimInterfaces = append(resultPimInterfaces, data.PimInterfaces[ni])
			}
		}
		data.PimInterfaces = resultPimInterfaces
	}
	oldStaticRpAddresses := data.StaticRpAddresses
	if value := res.Get(path + "pim.rpAddr"); value.Exists() && len(value.Array()) > 0 {
		data.StaticRpAddresses = make([]ServiceMulticastStaticRpAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceMulticastStaticRpAddresses{}
			item.IpAddress = types.StringNull()
			item.IpAddressVariable = types.StringNull()
			if t := v.Get("address.optionType"); t.Exists() {
				va := v.Get("address.value")
				if t.String() == "variable" {
					item.IpAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IpAddress = types.StringValue(va.String())
				}
			}
			item.AccessList = types.StringNull()
			item.AccessListVariable = types.StringNull()
			if t := v.Get("accessList.optionType"); t.Exists() {
				va := v.Get("accessList.value")
				if t.String() == "variable" {
					item.AccessListVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AccessList = types.StringValue(va.String())
				}
			}
			item.Override = types.BoolNull()
			item.OverrideVariable = types.StringNull()
			if t := v.Get("override.optionType"); t.Exists() {
				va := v.Get("override.value")
				if t.String() == "variable" {
					item.OverrideVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Override = types.BoolValue(va.Bool())
				}
			}
			data.StaticRpAddresses = append(data.StaticRpAddresses, item)
			return true
		})
	} else {
		data.StaticRpAddresses = nil
	}
	if !fullRead {
		resultStaticRpAddresses := make([]ServiceMulticastStaticRpAddresses, 0, len(data.StaticRpAddresses))
		matchedStaticRpAddresses := make([]bool, len(data.StaticRpAddresses))
		for _, oldItem := range oldStaticRpAddresses {
			for ni := range data.StaticRpAddresses {
				if matchedStaticRpAddresses[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.IpAddressVariable.ValueString() != "" || data.StaticRpAddresses[ni].IpAddressVariable.ValueString() != "") {
					if oldItem.IpAddressVariable.ValueString() != data.StaticRpAddresses[ni].IpAddressVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.IpAddress.ValueString() != data.StaticRpAddresses[ni].IpAddress.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedStaticRpAddresses[ni] = true
					resultStaticRpAddresses = append(resultStaticRpAddresses, data.StaticRpAddresses[ni])
					break
				}
			}
		}
		for ni := range data.StaticRpAddresses {
			if !matchedStaticRpAddresses[ni] {
				resultStaticRpAddresses = append(resultStaticRpAddresses, data.StaticRpAddresses[ni])
			}
		}
		data.StaticRpAddresses = resultStaticRpAddresses
	}
	data.EnableAutoRp = types.BoolNull()
	data.EnableAutoRpVariable = types.StringNull()
	if t := res.Get(path + "pim.autoRp.enableAutoRPFlag.optionType"); t.Exists() {
		va := res.Get(path + "pim.autoRp.enableAutoRPFlag.value")
		if t.String() == "variable" {
			data.EnableAutoRpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnableAutoRp = types.BoolValue(va.Bool())
		}
	}
	oldAutoRpAnnounces := data.AutoRpAnnounces
	if value := res.Get(path + "pim.autoRp.sendRpAnnounceList"); value.Exists() && len(value.Array()) > 0 {
		data.AutoRpAnnounces = make([]ServiceMulticastAutoRpAnnounces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceMulticastAutoRpAnnounces{}
			item.InterfaceName = types.StringNull()
			item.InterfaceNameVariable = types.StringNull()
			if t := v.Get("interfaceName.optionType"); t.Exists() {
				va := v.Get("interfaceName.value")
				if t.String() == "variable" {
					item.InterfaceNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.InterfaceName = types.StringValue(va.String())
				}
			}
			item.Scope = types.Int64Null()
			item.ScopeVariable = types.StringNull()
			if t := v.Get("scope.optionType"); t.Exists() {
				va := v.Get("scope.value")
				if t.String() == "variable" {
					item.ScopeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Scope = types.Int64Value(va.Int())
				}
			}
			data.AutoRpAnnounces = append(data.AutoRpAnnounces, item)
			return true
		})
	} else {
		data.AutoRpAnnounces = nil
	}
	if !fullRead {
		resultAutoRpAnnounces := make([]ServiceMulticastAutoRpAnnounces, 0, len(data.AutoRpAnnounces))
		matchedAutoRpAnnounces := make([]bool, len(data.AutoRpAnnounces))
		for _, oldItem := range oldAutoRpAnnounces {
			for ni := range data.AutoRpAnnounces {
				if matchedAutoRpAnnounces[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.InterfaceNameVariable.ValueString() != "" || data.AutoRpAnnounces[ni].InterfaceNameVariable.ValueString() != "") {
					if oldItem.InterfaceNameVariable.ValueString() != data.AutoRpAnnounces[ni].InterfaceNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.InterfaceName.ValueString() != data.AutoRpAnnounces[ni].InterfaceName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedAutoRpAnnounces[ni] = true
					resultAutoRpAnnounces = append(resultAutoRpAnnounces, data.AutoRpAnnounces[ni])
					break
				}
			}
		}
		for ni := range data.AutoRpAnnounces {
			if !matchedAutoRpAnnounces[ni] {
				resultAutoRpAnnounces = append(resultAutoRpAnnounces, data.AutoRpAnnounces[ni])
			}
		}
		data.AutoRpAnnounces = resultAutoRpAnnounces
	}
	oldAutoRpDiscoveries := data.AutoRpDiscoveries
	if value := res.Get(path + "pim.autoRp.sendRpDiscovery"); value.Exists() && len(value.Array()) > 0 {
		data.AutoRpDiscoveries = make([]ServiceMulticastAutoRpDiscoveries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceMulticastAutoRpDiscoveries{}
			item.InterfaceName = types.StringNull()
			item.InterfaceNameVariable = types.StringNull()
			if t := v.Get("interfaceName.optionType"); t.Exists() {
				va := v.Get("interfaceName.value")
				if t.String() == "variable" {
					item.InterfaceNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.InterfaceName = types.StringValue(va.String())
				}
			}
			item.Scope = types.Int64Null()
			item.ScopeVariable = types.StringNull()
			if t := v.Get("scope.optionType"); t.Exists() {
				va := v.Get("scope.value")
				if t.String() == "variable" {
					item.ScopeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Scope = types.Int64Value(va.Int())
				}
			}
			data.AutoRpDiscoveries = append(data.AutoRpDiscoveries, item)
			return true
		})
	} else {
		data.AutoRpDiscoveries = nil
	}
	if !fullRead {
		resultAutoRpDiscoveries := make([]ServiceMulticastAutoRpDiscoveries, 0, len(data.AutoRpDiscoveries))
		matchedAutoRpDiscoveries := make([]bool, len(data.AutoRpDiscoveries))
		for _, oldItem := range oldAutoRpDiscoveries {
			for ni := range data.AutoRpDiscoveries {
				if matchedAutoRpDiscoveries[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.InterfaceNameVariable.ValueString() != "" || data.AutoRpDiscoveries[ni].InterfaceNameVariable.ValueString() != "") {
					if oldItem.InterfaceNameVariable.ValueString() != data.AutoRpDiscoveries[ni].InterfaceNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.InterfaceName.ValueString() != data.AutoRpDiscoveries[ni].InterfaceName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedAutoRpDiscoveries[ni] = true
					resultAutoRpDiscoveries = append(resultAutoRpDiscoveries, data.AutoRpDiscoveries[ni])
					break
				}
			}
		}
		for ni := range data.AutoRpDiscoveries {
			if !matchedAutoRpDiscoveries[ni] {
				resultAutoRpDiscoveries = append(resultAutoRpDiscoveries, data.AutoRpDiscoveries[ni])
			}
		}
		data.AutoRpDiscoveries = resultAutoRpDiscoveries
	}
	oldPimBsrRpCandidates := data.PimBsrRpCandidates
	if value := res.Get(path + "pim.pimBsr.rpCandidate"); value.Exists() && len(value.Array()) > 0 {
		data.PimBsrRpCandidates = make([]ServiceMulticastPimBsrRpCandidates, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceMulticastPimBsrRpCandidates{}
			item.InterfaceName = types.StringNull()
			item.InterfaceNameVariable = types.StringNull()
			if t := v.Get("interfaceName.optionType"); t.Exists() {
				va := v.Get("interfaceName.value")
				if t.String() == "variable" {
					item.InterfaceNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.InterfaceName = types.StringValue(va.String())
				}
			}
			item.AccessListId = types.StringNull()
			item.AccessListIdVariable = types.StringNull()
			if t := v.Get("groupList.optionType"); t.Exists() {
				va := v.Get("groupList.value")
				if t.String() == "variable" {
					item.AccessListIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AccessListId = types.StringValue(va.String())
				}
			}
			item.Interval = types.Int64Null()
			item.IntervalVariable = types.StringNull()
			if t := v.Get("interval.optionType"); t.Exists() {
				va := v.Get("interval.value")
				if t.String() == "variable" {
					item.IntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Interval = types.Int64Value(va.Int())
				}
			}
			item.Priority = types.Int64Null()
			item.PriorityVariable = types.StringNull()
			if t := v.Get("priority.optionType"); t.Exists() {
				va := v.Get("priority.value")
				if t.String() == "variable" {
					item.PriorityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Priority = types.Int64Value(va.Int())
				}
			}
			data.PimBsrRpCandidates = append(data.PimBsrRpCandidates, item)
			return true
		})
	} else {
		data.PimBsrRpCandidates = nil
	}
	if !fullRead {
		resultPimBsrRpCandidates := make([]ServiceMulticastPimBsrRpCandidates, 0, len(data.PimBsrRpCandidates))
		matchedPimBsrRpCandidates := make([]bool, len(data.PimBsrRpCandidates))
		for _, oldItem := range oldPimBsrRpCandidates {
			for ni := range data.PimBsrRpCandidates {
				if matchedPimBsrRpCandidates[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.InterfaceNameVariable.ValueString() != "" || data.PimBsrRpCandidates[ni].InterfaceNameVariable.ValueString() != "") {
					if oldItem.InterfaceNameVariable.ValueString() != data.PimBsrRpCandidates[ni].InterfaceNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.InterfaceName.ValueString() != data.PimBsrRpCandidates[ni].InterfaceName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedPimBsrRpCandidates[ni] = true
					resultPimBsrRpCandidates = append(resultPimBsrRpCandidates, data.PimBsrRpCandidates[ni])
					break
				}
			}
		}
		for ni := range data.PimBsrRpCandidates {
			if !matchedPimBsrRpCandidates[ni] {
				resultPimBsrRpCandidates = append(resultPimBsrRpCandidates, data.PimBsrRpCandidates[ni])
			}
		}
		data.PimBsrRpCandidates = resultPimBsrRpCandidates
	}
	oldPimBsrCandidates := data.PimBsrCandidates
	if value := res.Get(path + "pim.pimBsr.bsrCandidate"); value.Exists() && len(value.Array()) > 0 {
		data.PimBsrCandidates = make([]ServiceMulticastPimBsrCandidates, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceMulticastPimBsrCandidates{}
			item.InterfaceName = types.StringNull()
			item.InterfaceNameVariable = types.StringNull()
			if t := v.Get("interfaceName.optionType"); t.Exists() {
				va := v.Get("interfaceName.value")
				if t.String() == "variable" {
					item.InterfaceNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.InterfaceName = types.StringValue(va.String())
				}
			}
			item.HashMaskLength = types.Int64Null()
			item.HashMaskLengthVariable = types.StringNull()
			if t := v.Get("mask.optionType"); t.Exists() {
				va := v.Get("mask.value")
				if t.String() == "variable" {
					item.HashMaskLengthVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.HashMaskLength = types.Int64Value(va.Int())
				}
			}
			item.Priority = types.Int64Null()
			item.PriorityVariable = types.StringNull()
			if t := v.Get("priority.optionType"); t.Exists() {
				va := v.Get("priority.value")
				if t.String() == "variable" {
					item.PriorityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Priority = types.Int64Value(va.Int())
				}
			}
			item.AcceptCandidateAccessList = types.StringNull()
			item.AcceptCandidateAccessListVariable = types.StringNull()
			if t := v.Get("acceptRpCandidate.optionType"); t.Exists() {
				va := v.Get("acceptRpCandidate.value")
				if t.String() == "variable" {
					item.AcceptCandidateAccessListVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AcceptCandidateAccessList = types.StringValue(va.String())
				}
			}
			data.PimBsrCandidates = append(data.PimBsrCandidates, item)
			return true
		})
	} else {
		data.PimBsrCandidates = nil
	}
	if !fullRead {
		resultPimBsrCandidates := make([]ServiceMulticastPimBsrCandidates, 0, len(data.PimBsrCandidates))
		matchedPimBsrCandidates := make([]bool, len(data.PimBsrCandidates))
		for _, oldItem := range oldPimBsrCandidates {
			for ni := range data.PimBsrCandidates {
				if matchedPimBsrCandidates[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.InterfaceNameVariable.ValueString() != "" || data.PimBsrCandidates[ni].InterfaceNameVariable.ValueString() != "") {
					if oldItem.InterfaceNameVariable.ValueString() != data.PimBsrCandidates[ni].InterfaceNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.InterfaceName.ValueString() != data.PimBsrCandidates[ni].InterfaceName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedPimBsrCandidates[ni] = true
					resultPimBsrCandidates = append(resultPimBsrCandidates, data.PimBsrCandidates[ni])
					break
				}
			}
		}
		for ni := range data.PimBsrCandidates {
			if !matchedPimBsrCandidates[ni] {
				resultPimBsrCandidates = append(resultPimBsrCandidates, data.PimBsrCandidates[ni])
			}
		}
		data.PimBsrCandidates = resultPimBsrCandidates
	}
	oldMsdpGroups := data.MsdpGroups
	if value := res.Get(path + "msdp.msdpList"); value.Exists() && len(value.Array()) > 0 {
		data.MsdpGroups = make([]ServiceMulticastMsdpGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceMulticastMsdpGroups{}
			item.MeshGroupName = types.StringNull()
			item.MeshGroupNameVariable = types.StringNull()
			if t := v.Get("meshGroup.optionType"); t.Exists() {
				va := v.Get("meshGroup.value")
				if t.String() == "variable" {
					item.MeshGroupNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.MeshGroupName = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("peer"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Peers = make([]ServiceMulticastMsdpGroupsPeers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceMulticastMsdpGroupsPeers{}
					cItem.PeerIp = types.StringNull()
					cItem.PeerIpVariable = types.StringNull()
					if t := cv.Get("peerIp.optionType"); t.Exists() {
						va := cv.Get("peerIp.value")
						if t.String() == "variable" {
							cItem.PeerIpVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.PeerIp = types.StringValue(va.String())
						}
					}
					cItem.ConnectionSourceInterface = types.StringNull()
					cItem.ConnectionSourceInterfaceVariable = types.StringNull()
					if t := cv.Get("connectSourceIntf.optionType"); t.Exists() {
						va := cv.Get("connectSourceIntf.value")
						if t.String() == "variable" {
							cItem.ConnectionSourceInterfaceVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.ConnectionSourceInterface = types.StringValue(va.String())
						}
					}
					cItem.RemoteAs = types.Int64Null()
					cItem.RemoteAsVariable = types.StringNull()
					if t := cv.Get("remoteAs.optionType"); t.Exists() {
						va := cv.Get("remoteAs.value")
						if t.String() == "variable" {
							cItem.RemoteAsVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.RemoteAs = types.Int64Value(va.Int())
						}
					}
					cItem.KeepaliveInterval = types.Int64Null()
					cItem.KeepaliveIntervalVariable = types.StringNull()
					if t := cv.Get("keepaliveInterval.optionType"); t.Exists() {
						va := cv.Get("keepaliveInterval.value")
						if t.String() == "variable" {
							cItem.KeepaliveIntervalVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.KeepaliveInterval = types.Int64Value(va.Int())
						}
					}
					cItem.KeepaliveHoldTime = types.Int64Null()
					cItem.KeepaliveHoldTimeVariable = types.StringNull()
					if t := cv.Get("keepaliveHoldTime.optionType"); t.Exists() {
						va := cv.Get("keepaliveHoldTime.value")
						if t.String() == "variable" {
							cItem.KeepaliveHoldTimeVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.KeepaliveHoldTime = types.Int64Value(va.Int())
						}
					}
					cItem.SaLimit = types.Int64Null()
					cItem.SaLimitVariable = types.StringNull()
					if t := cv.Get("saLimit.optionType"); t.Exists() {
						va := cv.Get("saLimit.value")
						if t.String() == "variable" {
							cItem.SaLimitVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.SaLimit = types.Int64Value(va.Int())
						}
					}
					cItem.DefaultPeer = types.BoolNull()

					if t := cv.Get("default.defaultPeer.optionType"); t.Exists() {
						va := cv.Get("default.defaultPeer.value")
						if t.String() == "global" {
							cItem.DefaultPeer = types.BoolValue(va.Bool())
						}
					}
					cItem.PrefixListId = types.StringNull()

					if t := cv.Get("default.prefixList.refId.optionType"); t.Exists() {
						va := cv.Get("default.prefixList.refId.value")
						if t.String() == "global" {
							cItem.PrefixListId = types.StringValue(va.String())
						}
					}
					item.Peers = append(item.Peers, cItem)
					return true
				})
			}
			data.MsdpGroups = append(data.MsdpGroups, item)
			return true
		})
	} else {
		data.MsdpGroups = nil
	}
	if !fullRead {
		resultMsdpGroups := make([]ServiceMulticastMsdpGroups, 0, len(data.MsdpGroups))
		matchedMsdpGroups := make([]bool, len(data.MsdpGroups))
		for _, oldItem := range oldMsdpGroups {
			for ni := range data.MsdpGroups {
				if matchedMsdpGroups[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.MeshGroupNameVariable.ValueString() != "" || data.MsdpGroups[ni].MeshGroupNameVariable.ValueString() != "") {
					if oldItem.MeshGroupNameVariable.ValueString() != data.MsdpGroups[ni].MeshGroupNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.MeshGroupName.ValueString() != data.MsdpGroups[ni].MeshGroupName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedMsdpGroups[ni] = true
					{
						resultC := make([]ServiceMulticastMsdpGroupsPeers, 0, len(data.MsdpGroups[ni].Peers))
						matchedC := make([]bool, len(data.MsdpGroups[ni].Peers))
						for _, oldCItem := range oldItem.Peers {
							for nci := range data.MsdpGroups[ni].Peers {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.PeerIpVariable.ValueString() != "" || data.MsdpGroups[ni].Peers[nci].PeerIpVariable.ValueString() != "") {
									if oldCItem.PeerIpVariable.ValueString() != data.MsdpGroups[ni].Peers[nci].PeerIpVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.PeerIp.ValueString() != data.MsdpGroups[ni].Peers[nci].PeerIp.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.MsdpGroups[ni].Peers[nci])
									break
								}
							}
						}
						for nci := range data.MsdpGroups[ni].Peers {
							if !matchedC[nci] {
								resultC = append(resultC, data.MsdpGroups[ni].Peers[nci])
							}
						}
						data.MsdpGroups[ni].Peers = resultC
					}
					resultMsdpGroups = append(resultMsdpGroups, data.MsdpGroups[ni])
					break
				}
			}
		}
		for ni := range data.MsdpGroups {
			if !matchedMsdpGroups[ni] {
				resultMsdpGroups = append(resultMsdpGroups, data.MsdpGroups[ni])
			}
		}
		data.MsdpGroups = resultMsdpGroups
	}
	data.MsdpOriginatorId = types.StringNull()
	data.MsdpOriginatorIdVariable = types.StringNull()
	if t := res.Get(path + "msdp.originatorId.optionType"); t.Exists() {
		va := res.Get(path + "msdp.originatorId.value")
		if t.String() == "variable" {
			data.MsdpOriginatorIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MsdpOriginatorId = types.StringValue(va.String())
		}
	}
	data.MsdpConnectionRetryInterval = types.Int64Null()
	data.MsdpConnectionRetryIntervalVariable = types.StringNull()
	if t := res.Get(path + "msdp.refreshTimer.optionType"); t.Exists() {
		va := res.Get(path + "msdp.refreshTimer.value")
		if t.String() == "variable" {
			data.MsdpConnectionRetryIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MsdpConnectionRetryInterval = types.Int64Value(va.Int())
		}
	}
}

// End of section. //template:end fromBody
