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
type ApplicationPriorityTrafficPolicy struct {
	Id               types.String                                `tfsdk:"id"`
	Version          types.Int64                                 `tfsdk:"version"`
	Name             types.String                                `tfsdk:"name"`
	Description      types.String                                `tfsdk:"description"`
	FeatureProfileId types.String                                `tfsdk:"feature_profile_id"`
	DefaultAction    types.String                                `tfsdk:"default_action"`
	Vpns             types.Set                                   `tfsdk:"vpns"`
	Direction        types.String                                `tfsdk:"direction"`
	Sequences        []ApplicationPriorityTrafficPolicySequences `tfsdk:"sequences"`
}

type ApplicationPriorityTrafficPolicySequences struct {
	SequenceId   types.Int64                                             `tfsdk:"sequence_id"`
	SequenceName types.String                                            `tfsdk:"sequence_name"`
	BaseAction   types.String                                            `tfsdk:"base_action"`
	Protocol     types.String                                            `tfsdk:"protocol"`
	MatchEntries []ApplicationPriorityTrafficPolicySequencesMatchEntries `tfsdk:"match_entries"`
	Actions      []ApplicationPriorityTrafficPolicySequencesActions      `tfsdk:"actions"`
}

type ApplicationPriorityTrafficPolicySequencesMatchEntries struct {
	ApplicationListId               types.String `tfsdk:"application_list_id"`
	SaasApplicationListId           types.String `tfsdk:"saas_application_list_id"`
	ServiceAreas                    types.Set    `tfsdk:"service_areas"`
	TrafficCategory                 types.String `tfsdk:"traffic_category"`
	DnsApplicationListId            types.String `tfsdk:"dns_application_list_id"`
	TrafficClass                    types.String `tfsdk:"traffic_class"`
	Dscps                           types.Set    `tfsdk:"dscps"`
	PacketLength                    types.String `tfsdk:"packet_length"`
	Protocols                       types.Set    `tfsdk:"protocols"`
	IcmpMessages                    types.Set    `tfsdk:"icmp_messages"`
	Icmp6Messages                   types.Set    `tfsdk:"icmp6_messages"`
	SourceDataIpv4PrefixListId      types.String `tfsdk:"source_data_ipv4_prefix_list_id"`
	SourceDataIpv6PrefixListId      types.String `tfsdk:"source_data_ipv6_prefix_list_id"`
	SourceIpv4Prefix                types.String `tfsdk:"source_ipv4_prefix"`
	SourceIpv6Prefix                types.String `tfsdk:"source_ipv6_prefix"`
	SourcePorts                     types.Set    `tfsdk:"source_ports"`
	DestinationDataIpv4PrefixListId types.String `tfsdk:"destination_data_ipv4_prefix_list_id"`
	DestinationDataIpv6PrefixListId types.String `tfsdk:"destination_data_ipv6_prefix_list_id"`
	DestinationIpv4Prefix           types.String `tfsdk:"destination_ipv4_prefix"`
	DestinationIpv6Prefix           types.String `tfsdk:"destination_ipv6_prefix"`
	DestinationPorts                types.Set    `tfsdk:"destination_ports"`
	Tcp                             types.String `tfsdk:"tcp"`
	DestinationRegion               types.String `tfsdk:"destination_region"`
	TrafficTo                       types.String `tfsdk:"traffic_to"`
	Dns                             types.String `tfsdk:"dns"`
}
type ApplicationPriorityTrafficPolicySequencesActions struct {
	SlaClasses                []ApplicationPriorityTrafficPolicySequencesActionsSlaClasses    `tfsdk:"sla_classes"`
	BackupSlaPreferredColors  types.Set                                                       `tfsdk:"backup_sla_preferred_colors"`
	SetParameters             []ApplicationPriorityTrafficPolicySequencesActionsSetParameters `tfsdk:"set_parameters"`
	RedirectDnsField          types.String                                                    `tfsdk:"redirect_dns_field"`
	RedirectDnsValue          types.String                                                    `tfsdk:"redirect_dns_value"`
	AppqoeTcpOptimization     types.Bool                                                      `tfsdk:"appqoe_tcp_optimization"`
	AppqoeDreOptimization     types.Bool                                                      `tfsdk:"appqoe_dre_optimization"`
	AppqoeServiceNodeGroup    types.String                                                    `tfsdk:"appqoe_service_node_group"`
	LossCorrectType           types.String                                                    `tfsdk:"loss_correct_type"`
	LossCorrectFecThreshold   types.Int64                                                     `tfsdk:"loss_correct_fec_threshold"`
	Count                     types.String                                                    `tfsdk:"count"`
	Log                       types.Bool                                                      `tfsdk:"log"`
	CloudSaas                 types.Bool                                                      `tfsdk:"cloud_saas"`
	CloudProbe                types.Bool                                                      `tfsdk:"cloud_probe"`
	Cflowd                    types.Bool                                                      `tfsdk:"cflowd"`
	NatPool                   types.Int64                                                     `tfsdk:"nat_pool"`
	NatVpn                    types.Bool                                                      `tfsdk:"nat_vpn"`
	NatFallback               types.Bool                                                      `tfsdk:"nat_fallback"`
	NatBypass                 types.Bool                                                      `tfsdk:"nat_bypass"`
	NatDiaPools               types.Set                                                       `tfsdk:"nat_dia_pools"`
	NatDiaInterfaces          types.Set                                                       `tfsdk:"nat_dia_interfaces"`
	SecureInternetGateway     types.Bool                                                      `tfsdk:"secure_internet_gateway"`
	FallbackToRouting         types.Bool                                                      `tfsdk:"fallback_to_routing"`
	SecureServiceEdge         types.Bool                                                      `tfsdk:"secure_service_edge"`
	SecureServiceEdgeInstance types.String                                                    `tfsdk:"secure_service_edge_instance"`
}

type ApplicationPriorityTrafficPolicySequencesActionsSlaClasses struct {
	SlaClassListId            types.String `tfsdk:"sla_class_list_id"`
	PreferredColors           types.Set    `tfsdk:"preferred_colors"`
	PreferredColorGroupListId types.String `tfsdk:"preferred_color_group_list_id"`
	Strict                    types.Bool   `tfsdk:"strict"`
	FallbackToBestPath        types.Bool   `tfsdk:"fallback_to_best_path"`
	PreferredRemoteColors     types.Set    `tfsdk:"preferred_remote_colors"`
	RemoteColorRestrict       types.Bool   `tfsdk:"remote_color_restrict"`
}
type ApplicationPriorityTrafficPolicySequencesActionsSetParameters struct {
	Dscp                          types.Int64  `tfsdk:"dscp"`
	PolicerId                     types.String `tfsdk:"policer_id"`
	PreferredColorGroupId         types.String `tfsdk:"preferred_color_group_id"`
	ForwardingClassListId         types.String `tfsdk:"forwarding_class_list_id"`
	LocalTlocListColors           types.Set    `tfsdk:"local_tloc_list_colors"`
	LocalTlocListRestrict         types.Bool   `tfsdk:"local_tloc_list_restrict"`
	LocalTlocListEncapsulation    types.String `tfsdk:"local_tloc_list_encapsulation"`
	PreferredRemoteColors         types.Set    `tfsdk:"preferred_remote_colors"`
	PreferredRemoteColorRestrict  types.Bool   `tfsdk:"preferred_remote_color_restrict"`
	TlocColor                     types.Set    `tfsdk:"tloc_color"`
	TlocEncapsulation             types.String `tfsdk:"tloc_encapsulation"`
	TlocIp                        types.String `tfsdk:"tloc_ip"`
	TlocListId                    types.String `tfsdk:"tloc_list_id"`
	ServiceTlocColor              types.Set    `tfsdk:"service_tloc_color"`
	ServiceTlocEncapsulation      types.String `tfsdk:"service_tloc_encapsulation"`
	ServiceTlocIp                 types.String `tfsdk:"service_tloc_ip"`
	ServiceVpn                    types.Int64  `tfsdk:"service_vpn"`
	ServiceType                   types.String `tfsdk:"service_type"`
	ServiceLocal                  types.Bool   `tfsdk:"service_local"`
	ServiceRestrict               types.Bool   `tfsdk:"service_restrict"`
	ServiceTlocListId             types.String `tfsdk:"service_tloc_list_id"`
	ServiceChainType              types.String `tfsdk:"service_chain_type"`
	ServiceChainVpn               types.Int64  `tfsdk:"service_chain_vpn"`
	ServiceChainLocal             types.Bool   `tfsdk:"service_chain_local"`
	ServiceChainFallbackToRouting types.Bool   `tfsdk:"service_chain_fallback_to_routing"`
	ServiceChainTlocColor         types.Set    `tfsdk:"service_chain_tloc_color"`
	ServiceChainTlocEncapsulation types.String `tfsdk:"service_chain_tloc_encapsulation"`
	ServiceChainTlocIp            types.String `tfsdk:"service_chain_tloc_ip"`
	ServiceChainTlocListId        types.String `tfsdk:"service_chain_tloc_list_id"`
	NextHopIpv4                   types.String `tfsdk:"next_hop_ipv4"`
	NextHopIpv6                   types.String `tfsdk:"next_hop_ipv6"`
	NextHopLoose                  types.Bool   `tfsdk:"next_hop_loose"`
	Vpn                           types.Int64  `tfsdk:"vpn"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ApplicationPriorityTrafficPolicy) getModel() string {
	return "application_priority_traffic_policy"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ApplicationPriorityTrafficPolicy) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/application-priority/%v/traffic-policy", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ApplicationPriorityTrafficPolicy) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.DefaultAction.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dataDefaultAction.optionType", "global")
			body, _ = sjson.Set(body, path+"dataDefaultAction.value", data.DefaultAction.ValueString())
		}
	}
	if !data.Vpns.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.vpn.optionType", "global")
			var values []string
			data.Vpns.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"target.vpn.value", values)
		}
	}
	if !data.Direction.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.direction.optionType", "global")
			body, _ = sjson.Set(body, path+"target.direction.value", data.Direction.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.SequenceId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceId.value", item.SequenceId.ValueInt64())
				}
			}
			if !item.SequenceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceName.value", item.SequenceName.ValueString())
				}
			}
			if !item.BaseAction.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "baseAction.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "baseAction.value", item.BaseAction.ValueString())
				}
			}
			if !item.Protocol.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceIpType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceIpType.value", item.Protocol.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.ApplicationListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "appList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "appList.refId.value", childItem.ApplicationListId.ValueString())
						}
					}
					if !childItem.SaasApplicationListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "saasAppList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "saasAppList.refId.value", childItem.SaasApplicationListId.ValueString())
						}
					}
					if !childItem.ServiceAreas.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "serviceArea.optionType", "global")
							var values []string
							childItem.ServiceAreas.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "serviceArea.value", values)
						}
					}
					if !childItem.TrafficCategory.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "trafficCategory.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "trafficCategory.value", childItem.TrafficCategory.ValueString())
						}
					}
					if !childItem.DnsApplicationListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "dnsAppList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "dnsAppList.refId.value", childItem.DnsApplicationListId.ValueString())
						}
					}
					if !childItem.TrafficClass.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "trafficClass.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "trafficClass.value", childItem.TrafficClass.ValueString())
						}
					}
					if !childItem.Dscps.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "dscp.optionType", "global")
							var values []int64
							childItem.Dscps.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "dscp.value", values)
						}
					}
					if !childItem.PacketLength.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "packetLength.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "packetLength.value", childItem.PacketLength.ValueString())
						}
					}
					if !childItem.Protocols.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "global")
							var values []string
							childItem.Protocols.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", values)
						}
					}
					if !childItem.IcmpMessages.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "icmpMessage.optionType", "global")
							var values []string
							childItem.IcmpMessages.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "icmpMessage.value", values)
						}
					}
					if !childItem.Icmp6Messages.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "icmp6Message.optionType", "global")
							var values []string
							childItem.Icmp6Messages.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "icmp6Message.value", values)
						}
					}
					if !childItem.SourceDataIpv4PrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefixList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefixList.refId.value", childItem.SourceDataIpv4PrefixListId.ValueString())
						}
					}
					if !childItem.SourceDataIpv6PrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataIpv6PrefixList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataIpv6PrefixList.refId.value", childItem.SourceDataIpv6PrefixListId.ValueString())
						}
					}
					if !childItem.SourceIpv4Prefix.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIp.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIp.value", childItem.SourceIpv4Prefix.ValueString())
						}
					}
					if !childItem.SourceIpv6Prefix.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIpv6.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIpv6.value", childItem.SourceIpv6Prefix.ValueString())
						}
					}
					if !childItem.SourcePorts.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourcePort.optionType", "global")
							var values []string
							childItem.SourcePorts.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourcePort.value", values)
						}
					}
					if !childItem.DestinationDataIpv4PrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefixList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefixList.refId.value", childItem.DestinationDataIpv4PrefixListId.ValueString())
						}
					}
					if !childItem.DestinationDataIpv6PrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataIpv6PrefixList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataIpv6PrefixList.refId.value", childItem.DestinationDataIpv6PrefixListId.ValueString())
						}
					}
					if !childItem.DestinationIpv4Prefix.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationIp.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationIp.value", childItem.DestinationIpv4Prefix.ValueString())
						}
					}
					if !childItem.DestinationIpv6Prefix.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationIpv6.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationIpv6.value", childItem.DestinationIpv6Prefix.ValueString())
						}
					}
					if !childItem.DestinationPorts.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationPort.optionType", "global")
							var values []string
							childItem.DestinationPorts.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationPort.value", values)
						}
					}
					if !childItem.Tcp.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "tcp.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "tcp.value", childItem.Tcp.ValueString())
						}
					}
					if !childItem.DestinationRegion.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationRegion.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationRegion.value", childItem.DestinationRegion.ValueString())
						}
					}
					if !childItem.TrafficTo.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "trafficTo.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "trafficTo.value", childItem.TrafficTo.ValueString())
						}
					}
					if !childItem.Dns.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "dns.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "dns.value", childItem.Dns.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "actions", []interface{}{})
				for _, childItem := range item.Actions {
					itemChildBody := ""
					if true {

						for _, childChildItem := range childItem.SlaClasses {
							itemChildChildBody := ""
							if !childChildItem.SlaClassListId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "slaName.refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "slaName.refId.value", childChildItem.SlaClassListId.ValueString())
								}
							}
							if !childChildItem.PreferredColors.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredColor.optionType", "global")
									var values []string
									childChildItem.PreferredColors.ElementsAs(ctx, &values, false)
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredColor.value", values)
								}
							}
							if !childChildItem.PreferredColorGroupListId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredColorGroup.refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredColorGroup.refId.value", childChildItem.PreferredColorGroupListId.ValueString())
								}
							}
							if !childChildItem.Strict.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "strict.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "strict.value", childChildItem.Strict.ValueBool())
								}
							}
							if !childChildItem.FallbackToBestPath.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "fallbackToBestPath.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "fallbackToBestPath.value", childChildItem.FallbackToBestPath.ValueBool())
								}
							}
							if !childChildItem.PreferredRemoteColors.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredRemoteColor.optionType", "global")
									var values []string
									childChildItem.PreferredRemoteColors.ElementsAs(ctx, &values, false)
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredRemoteColor.value", values)
								}
							}
							if !childChildItem.RemoteColorRestrict.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "remoteColorRestrict.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "remoteColorRestrict.value", childChildItem.RemoteColorRestrict.ValueBool())
								}
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "slaClass.-1", itemChildChildBody)
						}
					}
					if !childItem.BackupSlaPreferredColors.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "backupSlaPreferredColor.optionType", "global")
							var values []string
							childItem.BackupSlaPreferredColors.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "backupSlaPreferredColor.value", values)
						}
					}
					if true {

						for _, childChildItem := range childItem.SetParameters {
							itemChildChildBody := ""
							if !childChildItem.Dscp.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "dscp.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "dscp.value", childChildItem.Dscp.ValueInt64())
								}
							}
							if !childChildItem.PolicerId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "policer.refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "policer.refId.value", childChildItem.PolicerId.ValueString())
								}
							}
							if !childChildItem.PreferredColorGroupId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredColorGroup.refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredColorGroup.refId.value", childChildItem.PreferredColorGroupId.ValueString())
								}
							}
							if !childChildItem.ForwardingClassListId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "forwardingClass.refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "forwardingClass.refId.value", childChildItem.ForwardingClassListId.ValueString())
								}
							}
							if !childChildItem.LocalTlocListColors.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.color.optionType", "global")
									var values []string
									childChildItem.LocalTlocListColors.ElementsAs(ctx, &values, false)
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.color.value", values)
								}
							}
							if !childChildItem.LocalTlocListRestrict.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.restrict.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.restrict.value", childChildItem.LocalTlocListRestrict.ValueBool())
								}
							}
							if !childChildItem.LocalTlocListEncapsulation.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.encap.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.encap.value", childChildItem.LocalTlocListEncapsulation.ValueString())
								}
							}
							if !childChildItem.PreferredRemoteColors.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredRemoteColor.color.optionType", "global")
									var values []string
									childChildItem.PreferredRemoteColors.ElementsAs(ctx, &values, false)
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredRemoteColor.color.value", values)
								}
							}
							if !childChildItem.PreferredRemoteColorRestrict.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredRemoteColor.remoteColorRestrict.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredRemoteColor.remoteColorRestrict.value", childChildItem.PreferredRemoteColorRestrict.ValueBool())
								}
							}
							if !childChildItem.TlocColor.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.color.optionType", "global")
									var values []string
									childChildItem.TlocColor.ElementsAs(ctx, &values, false)
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.color.value", values)
								}
							}
							if !childChildItem.TlocEncapsulation.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.encap.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.encap.value", childChildItem.TlocEncapsulation.ValueString())
								}
							}
							if !childChildItem.TlocIp.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.ip.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.ip.value", childChildItem.TlocIp.ValueString())
								}
							}
							if !childChildItem.TlocListId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocList.refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocList.refId.value", childChildItem.TlocListId.ValueString())
								}
							}
							if !childChildItem.ServiceTlocColor.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.color.optionType", "global")
									var values []string
									childChildItem.ServiceTlocColor.ElementsAs(ctx, &values, false)
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.color.value", values)
								}
							}
							if !childChildItem.ServiceTlocEncapsulation.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.encap.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.encap.value", childChildItem.ServiceTlocEncapsulation.ValueString())
								}
							}
							if !childChildItem.ServiceTlocIp.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.ip.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.ip.value", childChildItem.ServiceTlocIp.ValueString())
								}
							}
							if !childChildItem.ServiceVpn.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.vpn.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.vpn.value", childChildItem.ServiceVpn.ValueInt64())
								}
							}
							if !childChildItem.ServiceType.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.type.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.type.value", childChildItem.ServiceType.ValueString())
								}
							}
							if !childChildItem.ServiceLocal.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.local.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.local.value", childChildItem.ServiceLocal.ValueBool())
								}
							}
							if !childChildItem.ServiceRestrict.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.restrict.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.restrict.value", childChildItem.ServiceRestrict.ValueBool())
								}
							}
							if !childChildItem.ServiceTlocListId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tlocList.refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tlocList.refId.value", childChildItem.ServiceTlocListId.ValueString())
								}
							}
							if !childChildItem.ServiceChainType.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.type.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.type.value", childChildItem.ServiceChainType.ValueString())
								}
							}
							if !childChildItem.ServiceChainVpn.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.vpn.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.vpn.value", childChildItem.ServiceChainVpn.ValueInt64())
								}
							}
							if !childChildItem.ServiceChainLocal.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.local.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.local.value", childChildItem.ServiceChainLocal.ValueBool())
								}
							}
							if !childChildItem.ServiceChainFallbackToRouting.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.restrict.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.restrict.value", childChildItem.ServiceChainFallbackToRouting.ValueBool())
								}
							}
							if !childChildItem.ServiceChainTlocColor.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.color.optionType", "global")
									var values []string
									childChildItem.ServiceChainTlocColor.ElementsAs(ctx, &values, false)
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.color.value", values)
								}
							}
							if !childChildItem.ServiceChainTlocEncapsulation.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.encap.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.encap.value", childChildItem.ServiceChainTlocEncapsulation.ValueString())
								}
							}
							if !childChildItem.ServiceChainTlocIp.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.ip.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.ip.value", childChildItem.ServiceChainTlocIp.ValueString())
								}
							}
							if !childChildItem.ServiceChainTlocListId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tlocList.refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tlocList.refId.value", childChildItem.ServiceChainTlocListId.ValueString())
								}
							}
							if !childChildItem.NextHopIpv4.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "nextHop.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "nextHop.value", childChildItem.NextHopIpv4.ValueString())
								}
							}
							if !childChildItem.NextHopIpv6.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "nextHopIpv6.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "nextHopIpv6.value", childChildItem.NextHopIpv6.ValueString())
								}
							}
							if !childChildItem.NextHopLoose.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "nextHopLoose.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "nextHopLoose.value", childChildItem.NextHopLoose.ValueBool())
								}
							}
							if !childChildItem.Vpn.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "vpn.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "vpn.value", childChildItem.Vpn.ValueInt64())
								}
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "set.-1", itemChildChildBody)
						}
					}
					if !childItem.RedirectDnsField.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "redirectDns.field.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "redirectDns.field.value", childItem.RedirectDnsField.ValueString())
						}
					}
					if !childItem.RedirectDnsValue.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "redirectDns.value.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "redirectDns.value.value", childItem.RedirectDnsValue.ValueString())
						}
					}
					if !childItem.AppqoeTcpOptimization.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "appqoeOptimization.tcpOptimization.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "appqoeOptimization.tcpOptimization.value", childItem.AppqoeTcpOptimization.ValueBool())
						}
					}
					if !childItem.AppqoeDreOptimization.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "appqoeOptimization.dreOptimization.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "appqoeOptimization.dreOptimization.value", childItem.AppqoeDreOptimization.ValueBool())
						}
					}
					if !childItem.AppqoeServiceNodeGroup.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "appqoeOptimization.serviceNodeGroup.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "appqoeOptimization.serviceNodeGroup.value", childItem.AppqoeServiceNodeGroup.ValueString())
						}
					}
					if !childItem.LossCorrectType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "lossCorrection.lossCorrectionType.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "lossCorrection.lossCorrectionType.value", childItem.LossCorrectType.ValueString())
						}
					}
					if !childItem.LossCorrectFecThreshold.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "lossCorrection.lossCorrectFec.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "lossCorrection.lossCorrectFec.value", childItem.LossCorrectFecThreshold.ValueInt64())
						}
					}
					if !childItem.Count.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "count.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "count.value", childItem.Count.ValueString())
						}
					}
					if !childItem.Log.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "log.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "log.value", childItem.Log.ValueBool())
						}
					}
					if !childItem.CloudSaas.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "cloudSaas.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "cloudSaas.value", childItem.CloudSaas.ValueBool())
						}
					}
					if !childItem.CloudProbe.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "cloudProbe.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "cloudProbe.value", childItem.CloudProbe.ValueBool())
						}
					}
					if !childItem.Cflowd.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "cflowd.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "cflowd.value", childItem.Cflowd.ValueBool())
						}
					}
					if !childItem.NatPool.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "natPool.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "natPool.value", childItem.NatPool.ValueInt64())
						}
					}
					if !childItem.NatVpn.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.useVpn.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.useVpn.value", childItem.NatVpn.ValueBool())
						}
					}
					if !childItem.NatFallback.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.fallback.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.fallback.value", childItem.NatFallback.ValueBool())
						}
					}
					if !childItem.NatBypass.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.bypass.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.bypass.value", childItem.NatBypass.ValueBool())
						}
					}
					if !childItem.NatDiaPools.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.diaPool.optionType", "global")
							var values []int64
							childItem.NatDiaPools.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.diaPool.value", values)
						}
					}
					if !childItem.NatDiaInterfaces.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.diaInterface.optionType", "global")
							var values []string
							childItem.NatDiaInterfaces.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.diaInterface.value", values)
						}
					}
					if !childItem.SecureInternetGateway.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sig.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sig.value", childItem.SecureInternetGateway.ValueBool())
						}
					}
					if !childItem.FallbackToRouting.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "fallbackToRouting.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "fallbackToRouting.value", childItem.FallbackToRouting.ValueBool())
						}
					}
					if !childItem.SecureServiceEdge.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sse.secureServiceEdge.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sse.secureServiceEdge.value", childItem.SecureServiceEdge.ValueBool())
						}
					}
					if !childItem.SecureServiceEdgeInstance.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sse.secureServiceEdgeInstance.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sse.secureServiceEdgeInstance.value", childItem.SecureServiceEdgeInstance.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"sequences.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ApplicationPriorityTrafficPolicy) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "dataDefaultAction.optionType"); t.Exists() {
		va := res.Get(path + "dataDefaultAction.value")
		if t.String() == "global" {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	data.Vpns = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			data.Vpns = helpers.GetStringSet(va.Array())
		}
	}
	data.Direction = types.StringNull()

	if t := res.Get(path + "target.direction.optionType"); t.Exists() {
		va := res.Get(path + "target.direction.value")
		if t.String() == "global" {
			data.Direction = types.StringValue(va.String())
		}
	}
	oldSequences := data.Sequences
	if value := res.Get(path + "sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]ApplicationPriorityTrafficPolicySequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ApplicationPriorityTrafficPolicySequences{}
			item.SequenceId = types.Int64Null()

			if t := v.Get("sequenceId.optionType"); t.Exists() {
				va := v.Get("sequenceId.value")
				if t.String() == "global" {
					item.SequenceId = types.Int64Value(va.Int())
				}
			}
			item.SequenceName = types.StringNull()

			if t := v.Get("sequenceName.optionType"); t.Exists() {
				va := v.Get("sequenceName.value")
				if t.String() == "global" {
					item.SequenceName = types.StringValue(va.String())
				}
			}
			item.BaseAction = types.StringNull()

			if t := v.Get("baseAction.optionType"); t.Exists() {
				va := v.Get("baseAction.value")
				if t.String() == "global" {
					item.BaseAction = types.StringValue(va.String())
				}
			}
			item.Protocol = types.StringNull()

			if t := v.Get("sequenceIpType.optionType"); t.Exists() {
				va := v.Get("sequenceIpType.value")
				if t.String() == "global" {
					item.Protocol = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]ApplicationPriorityTrafficPolicySequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ApplicationPriorityTrafficPolicySequencesMatchEntries{}
					cItem.ApplicationListId = types.StringNull()

					if t := cv.Get("appList.refId.optionType"); t.Exists() {
						va := cv.Get("appList.refId.value")
						if t.String() == "global" {
							cItem.ApplicationListId = types.StringValue(va.String())
						}
					}
					cItem.SaasApplicationListId = types.StringNull()

					if t := cv.Get("saasAppList.refId.optionType"); t.Exists() {
						va := cv.Get("saasAppList.refId.value")
						if t.String() == "global" {
							cItem.SaasApplicationListId = types.StringValue(va.String())
						}
					}
					cItem.ServiceAreas = types.SetNull(types.StringType)

					if t := cv.Get("serviceArea.optionType"); t.Exists() {
						va := cv.Get("serviceArea.value")
						if t.String() == "global" {
							cItem.ServiceAreas = helpers.GetStringSet(va.Array())
						}
					}
					cItem.TrafficCategory = types.StringNull()

					if t := cv.Get("trafficCategory.optionType"); t.Exists() {
						va := cv.Get("trafficCategory.value")
						if t.String() == "global" {
							cItem.TrafficCategory = types.StringValue(va.String())
						}
					}
					cItem.DnsApplicationListId = types.StringNull()

					if t := cv.Get("dnsAppList.refId.optionType"); t.Exists() {
						va := cv.Get("dnsAppList.refId.value")
						if t.String() == "global" {
							cItem.DnsApplicationListId = types.StringValue(va.String())
						}
					}
					cItem.TrafficClass = types.StringNull()

					if t := cv.Get("trafficClass.optionType"); t.Exists() {
						va := cv.Get("trafficClass.value")
						if t.String() == "global" {
							cItem.TrafficClass = types.StringValue(va.String())
						}
					}
					cItem.Dscps = types.SetNull(types.Int64Type)

					if t := cv.Get("dscp.optionType"); t.Exists() {
						va := cv.Get("dscp.value")
						if t.String() == "global" {
							cItem.Dscps = helpers.GetInt64Set(va.Array())
						}
					}
					cItem.PacketLength = types.StringNull()

					if t := cv.Get("packetLength.optionType"); t.Exists() {
						va := cv.Get("packetLength.value")
						if t.String() == "global" {
							cItem.PacketLength = types.StringValue(va.String())
						}
					}
					cItem.Protocols = types.SetNull(types.StringType)

					if t := cv.Get("protocol.optionType"); t.Exists() {
						va := cv.Get("protocol.value")
						if t.String() == "global" {
							cItem.Protocols = helpers.GetStringSet(va.Array())
						}
					}
					cItem.IcmpMessages = types.SetNull(types.StringType)

					if t := cv.Get("icmpMessage.optionType"); t.Exists() {
						va := cv.Get("icmpMessage.value")
						if t.String() == "global" {
							cItem.IcmpMessages = helpers.GetStringSet(va.Array())
						}
					}
					cItem.Icmp6Messages = types.SetNull(types.StringType)

					if t := cv.Get("icmp6Message.optionType"); t.Exists() {
						va := cv.Get("icmp6Message.value")
						if t.String() == "global" {
							cItem.Icmp6Messages = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceDataIpv4PrefixListId = types.StringNull()

					if t := cv.Get("sourceDataPrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("sourceDataPrefixList.refId.value")
						if t.String() == "global" {
							cItem.SourceDataIpv4PrefixListId = types.StringValue(va.String())
						}
					}
					cItem.SourceDataIpv6PrefixListId = types.StringNull()

					if t := cv.Get("sourceDataIpv6PrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("sourceDataIpv6PrefixList.refId.value")
						if t.String() == "global" {
							cItem.SourceDataIpv6PrefixListId = types.StringValue(va.String())
						}
					}
					cItem.SourceIpv4Prefix = types.StringNull()

					if t := cv.Get("sourceIp.optionType"); t.Exists() {
						va := cv.Get("sourceIp.value")
						if t.String() == "global" {
							cItem.SourceIpv4Prefix = types.StringValue(va.String())
						}
					}
					cItem.SourceIpv6Prefix = types.StringNull()

					if t := cv.Get("sourceIpv6.optionType"); t.Exists() {
						va := cv.Get("sourceIpv6.value")
						if t.String() == "global" {
							cItem.SourceIpv6Prefix = types.StringValue(va.String())
						}
					}
					cItem.SourcePorts = types.SetNull(types.StringType)

					if t := cv.Get("sourcePort.optionType"); t.Exists() {
						va := cv.Get("sourcePort.value")
						if t.String() == "global" {
							cItem.SourcePorts = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationDataIpv4PrefixListId = types.StringNull()

					if t := cv.Get("destinationDataPrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("destinationDataPrefixList.refId.value")
						if t.String() == "global" {
							cItem.DestinationDataIpv4PrefixListId = types.StringValue(va.String())
						}
					}
					cItem.DestinationDataIpv6PrefixListId = types.StringNull()

					if t := cv.Get("destinationDataIpv6PrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("destinationDataIpv6PrefixList.refId.value")
						if t.String() == "global" {
							cItem.DestinationDataIpv6PrefixListId = types.StringValue(va.String())
						}
					}
					cItem.DestinationIpv4Prefix = types.StringNull()

					if t := cv.Get("destinationIp.optionType"); t.Exists() {
						va := cv.Get("destinationIp.value")
						if t.String() == "global" {
							cItem.DestinationIpv4Prefix = types.StringValue(va.String())
						}
					}
					cItem.DestinationIpv6Prefix = types.StringNull()

					if t := cv.Get("destinationIpv6.optionType"); t.Exists() {
						va := cv.Get("destinationIpv6.value")
						if t.String() == "global" {
							cItem.DestinationIpv6Prefix = types.StringValue(va.String())
						}
					}
					cItem.DestinationPorts = types.SetNull(types.StringType)

					if t := cv.Get("destinationPort.optionType"); t.Exists() {
						va := cv.Get("destinationPort.value")
						if t.String() == "global" {
							cItem.DestinationPorts = helpers.GetStringSet(va.Array())
						}
					}
					cItem.Tcp = types.StringNull()

					if t := cv.Get("tcp.optionType"); t.Exists() {
						va := cv.Get("tcp.value")
						if t.String() == "global" {
							cItem.Tcp = types.StringValue(va.String())
						}
					}
					cItem.DestinationRegion = types.StringNull()

					if t := cv.Get("destinationRegion.optionType"); t.Exists() {
						va := cv.Get("destinationRegion.value")
						if t.String() == "global" {
							cItem.DestinationRegion = types.StringValue(va.String())
						}
					}
					cItem.TrafficTo = types.StringNull()

					if t := cv.Get("trafficTo.optionType"); t.Exists() {
						va := cv.Get("trafficTo.value")
						if t.String() == "global" {
							cItem.TrafficTo = types.StringValue(va.String())
						}
					}
					cItem.Dns = types.StringNull()

					if t := cv.Get("dns.optionType"); t.Exists() {
						va := cv.Get("dns.value")
						if t.String() == "global" {
							cItem.Dns = types.StringValue(va.String())
						}
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Actions = make([]ApplicationPriorityTrafficPolicySequencesActions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ApplicationPriorityTrafficPolicySequencesActions{}
					if ccValue := cv.Get("slaClass"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.SlaClasses = make([]ApplicationPriorityTrafficPolicySequencesActionsSlaClasses, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := ApplicationPriorityTrafficPolicySequencesActionsSlaClasses{}
							ccItem.SlaClassListId = types.StringNull()

							if t := ccv.Get("slaName.refId.optionType"); t.Exists() {
								va := ccv.Get("slaName.refId.value")
								if t.String() == "global" {
									ccItem.SlaClassListId = types.StringValue(va.String())
								}
							}
							ccItem.PreferredColors = types.SetNull(types.StringType)

							if t := ccv.Get("preferredColor.optionType"); t.Exists() {
								va := ccv.Get("preferredColor.value")
								if t.String() == "global" {
									ccItem.PreferredColors = helpers.GetStringSet(va.Array())
								}
							}
							ccItem.PreferredColorGroupListId = types.StringNull()

							if t := ccv.Get("preferredColorGroup.refId.optionType"); t.Exists() {
								va := ccv.Get("preferredColorGroup.refId.value")
								if t.String() == "global" {
									ccItem.PreferredColorGroupListId = types.StringValue(va.String())
								}
							}
							ccItem.Strict = types.BoolNull()

							if t := ccv.Get("strict.optionType"); t.Exists() {
								va := ccv.Get("strict.value")
								if t.String() == "global" {
									ccItem.Strict = types.BoolValue(va.Bool())
								}
							}
							ccItem.FallbackToBestPath = types.BoolNull()

							if t := ccv.Get("fallbackToBestPath.optionType"); t.Exists() {
								va := ccv.Get("fallbackToBestPath.value")
								if t.String() == "global" {
									ccItem.FallbackToBestPath = types.BoolValue(va.Bool())
								}
							}
							ccItem.PreferredRemoteColors = types.SetNull(types.StringType)

							if t := ccv.Get("preferredRemoteColor.optionType"); t.Exists() {
								va := ccv.Get("preferredRemoteColor.value")
								if t.String() == "global" {
									ccItem.PreferredRemoteColors = helpers.GetStringSet(va.Array())
								}
							}
							ccItem.RemoteColorRestrict = types.BoolNull()

							if t := ccv.Get("remoteColorRestrict.optionType"); t.Exists() {
								va := ccv.Get("remoteColorRestrict.value")
								if t.String() == "global" {
									ccItem.RemoteColorRestrict = types.BoolValue(va.Bool())
								}
							}
							cItem.SlaClasses = append(cItem.SlaClasses, ccItem)
							return true
						})
					}
					cItem.BackupSlaPreferredColors = types.SetNull(types.StringType)

					if t := cv.Get("backupSlaPreferredColor.optionType"); t.Exists() {
						va := cv.Get("backupSlaPreferredColor.value")
						if t.String() == "global" {
							cItem.BackupSlaPreferredColors = helpers.GetStringSet(va.Array())
						}
					}
					if ccValue := cv.Get("set"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.SetParameters = make([]ApplicationPriorityTrafficPolicySequencesActionsSetParameters, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := ApplicationPriorityTrafficPolicySequencesActionsSetParameters{}
							ccItem.Dscp = types.Int64Null()

							if t := ccv.Get("dscp.optionType"); t.Exists() {
								va := ccv.Get("dscp.value")
								if t.String() == "global" {
									ccItem.Dscp = types.Int64Value(va.Int())
								}
							}
							ccItem.PolicerId = types.StringNull()

							if t := ccv.Get("policer.refId.optionType"); t.Exists() {
								va := ccv.Get("policer.refId.value")
								if t.String() == "global" {
									ccItem.PolicerId = types.StringValue(va.String())
								}
							}
							ccItem.PreferredColorGroupId = types.StringNull()

							if t := ccv.Get("preferredColorGroup.refId.optionType"); t.Exists() {
								va := ccv.Get("preferredColorGroup.refId.value")
								if t.String() == "global" {
									ccItem.PreferredColorGroupId = types.StringValue(va.String())
								}
							}
							ccItem.ForwardingClassListId = types.StringNull()

							if t := ccv.Get("forwardingClass.refId.optionType"); t.Exists() {
								va := ccv.Get("forwardingClass.refId.value")
								if t.String() == "global" {
									ccItem.ForwardingClassListId = types.StringValue(va.String())
								}
							}
							ccItem.LocalTlocListColors = types.SetNull(types.StringType)

							if t := ccv.Get("localTlocList.color.optionType"); t.Exists() {
								va := ccv.Get("localTlocList.color.value")
								if t.String() == "global" {
									ccItem.LocalTlocListColors = helpers.GetStringSet(va.Array())
								}
							}
							ccItem.LocalTlocListRestrict = types.BoolNull()

							if t := ccv.Get("localTlocList.restrict.optionType"); t.Exists() {
								va := ccv.Get("localTlocList.restrict.value")
								if t.String() == "global" {
									ccItem.LocalTlocListRestrict = types.BoolValue(va.Bool())
								}
							}
							ccItem.LocalTlocListEncapsulation = types.StringNull()

							if t := ccv.Get("localTlocList.encap.optionType"); t.Exists() {
								va := ccv.Get("localTlocList.encap.value")
								if t.String() == "global" {
									ccItem.LocalTlocListEncapsulation = types.StringValue(va.String())
								}
							}
							ccItem.PreferredRemoteColors = types.SetNull(types.StringType)

							if t := ccv.Get("preferredRemoteColor.color.optionType"); t.Exists() {
								va := ccv.Get("preferredRemoteColor.color.value")
								if t.String() == "global" {
									ccItem.PreferredRemoteColors = helpers.GetStringSet(va.Array())
								}
							}
							ccItem.PreferredRemoteColorRestrict = types.BoolNull()

							if t := ccv.Get("preferredRemoteColor.remoteColorRestrict.optionType"); t.Exists() {
								va := ccv.Get("preferredRemoteColor.remoteColorRestrict.value")
								if t.String() == "global" {
									ccItem.PreferredRemoteColorRestrict = types.BoolValue(va.Bool())
								}
							}
							ccItem.TlocColor = types.SetNull(types.StringType)

							if t := ccv.Get("tloc.color.optionType"); t.Exists() {
								va := ccv.Get("tloc.color.value")
								if t.String() == "global" {
									ccItem.TlocColor = helpers.GetStringSet(va.Array())
								}
							}
							ccItem.TlocEncapsulation = types.StringNull()

							if t := ccv.Get("tloc.encap.optionType"); t.Exists() {
								va := ccv.Get("tloc.encap.value")
								if t.String() == "global" {
									ccItem.TlocEncapsulation = types.StringValue(va.String())
								}
							}
							ccItem.TlocIp = types.StringNull()

							if t := ccv.Get("tloc.ip.optionType"); t.Exists() {
								va := ccv.Get("tloc.ip.value")
								if t.String() == "global" {
									ccItem.TlocIp = types.StringValue(va.String())
								}
							}
							ccItem.TlocListId = types.StringNull()

							if t := ccv.Get("tlocList.refId.optionType"); t.Exists() {
								va := ccv.Get("tlocList.refId.value")
								if t.String() == "global" {
									ccItem.TlocListId = types.StringValue(va.String())
								}
							}
							ccItem.ServiceTlocColor = types.SetNull(types.StringType)

							if t := ccv.Get("service.tloc.color.optionType"); t.Exists() {
								va := ccv.Get("service.tloc.color.value")
								if t.String() == "global" {
									ccItem.ServiceTlocColor = helpers.GetStringSet(va.Array())
								}
							}
							ccItem.ServiceTlocEncapsulation = types.StringNull()

							if t := ccv.Get("service.tloc.encap.optionType"); t.Exists() {
								va := ccv.Get("service.tloc.encap.value")
								if t.String() == "global" {
									ccItem.ServiceTlocEncapsulation = types.StringValue(va.String())
								}
							}
							ccItem.ServiceTlocIp = types.StringNull()

							if t := ccv.Get("service.tloc.ip.optionType"); t.Exists() {
								va := ccv.Get("service.tloc.ip.value")
								if t.String() == "global" {
									ccItem.ServiceTlocIp = types.StringValue(va.String())
								}
							}
							ccItem.ServiceVpn = types.Int64Null()

							if t := ccv.Get("service.vpn.optionType"); t.Exists() {
								va := ccv.Get("service.vpn.value")
								if t.String() == "global" {
									ccItem.ServiceVpn = types.Int64Value(va.Int())
								}
							}
							ccItem.ServiceType = types.StringNull()

							if t := ccv.Get("service.type.optionType"); t.Exists() {
								va := ccv.Get("service.type.value")
								if t.String() == "global" {
									ccItem.ServiceType = types.StringValue(va.String())
								}
							}
							ccItem.ServiceLocal = types.BoolNull()

							if t := ccv.Get("service.local.optionType"); t.Exists() {
								va := ccv.Get("service.local.value")
								if t.String() == "global" {
									ccItem.ServiceLocal = types.BoolValue(va.Bool())
								}
							}
							ccItem.ServiceRestrict = types.BoolNull()

							if t := ccv.Get("service.restrict.optionType"); t.Exists() {
								va := ccv.Get("service.restrict.value")
								if t.String() == "global" {
									ccItem.ServiceRestrict = types.BoolValue(va.Bool())
								}
							}
							ccItem.ServiceTlocListId = types.StringNull()

							if t := ccv.Get("service.tlocList.refId.optionType"); t.Exists() {
								va := ccv.Get("service.tlocList.refId.value")
								if t.String() == "global" {
									ccItem.ServiceTlocListId = types.StringValue(va.String())
								}
							}
							ccItem.ServiceChainType = types.StringNull()

							if t := ccv.Get("serviceChain.type.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.type.value")
								if t.String() == "global" {
									ccItem.ServiceChainType = types.StringValue(va.String())
								}
							}
							ccItem.ServiceChainVpn = types.Int64Null()

							if t := ccv.Get("serviceChain.vpn.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.vpn.value")
								if t.String() == "global" {
									ccItem.ServiceChainVpn = types.Int64Value(va.Int())
								}
							}
							ccItem.ServiceChainLocal = types.BoolNull()

							if t := ccv.Get("serviceChain.local.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.local.value")
								if t.String() == "global" {
									ccItem.ServiceChainLocal = types.BoolValue(va.Bool())
								}
							}
							ccItem.ServiceChainFallbackToRouting = types.BoolNull()

							if t := ccv.Get("serviceChain.restrict.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.restrict.value")
								if t.String() == "global" {
									ccItem.ServiceChainFallbackToRouting = types.BoolValue(va.Bool())
								}
							}
							ccItem.ServiceChainTlocColor = types.SetNull(types.StringType)

							if t := ccv.Get("serviceChain.tloc.color.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.tloc.color.value")
								if t.String() == "global" {
									ccItem.ServiceChainTlocColor = helpers.GetStringSet(va.Array())
								}
							}
							ccItem.ServiceChainTlocEncapsulation = types.StringNull()

							if t := ccv.Get("serviceChain.tloc.encap.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.tloc.encap.value")
								if t.String() == "global" {
									ccItem.ServiceChainTlocEncapsulation = types.StringValue(va.String())
								}
							}
							ccItem.ServiceChainTlocIp = types.StringNull()

							if t := ccv.Get("serviceChain.tloc.ip.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.tloc.ip.value")
								if t.String() == "global" {
									ccItem.ServiceChainTlocIp = types.StringValue(va.String())
								}
							}
							ccItem.ServiceChainTlocListId = types.StringNull()

							if t := ccv.Get("serviceChain.tlocList.refId.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.tlocList.refId.value")
								if t.String() == "global" {
									ccItem.ServiceChainTlocListId = types.StringValue(va.String())
								}
							}
							ccItem.NextHopIpv4 = types.StringNull()

							if t := ccv.Get("nextHop.optionType"); t.Exists() {
								va := ccv.Get("nextHop.value")
								if t.String() == "global" {
									ccItem.NextHopIpv4 = types.StringValue(va.String())
								}
							}
							ccItem.NextHopIpv6 = types.StringNull()

							if t := ccv.Get("nextHopIpv6.optionType"); t.Exists() {
								va := ccv.Get("nextHopIpv6.value")
								if t.String() == "global" {
									ccItem.NextHopIpv6 = types.StringValue(va.String())
								}
							}
							ccItem.NextHopLoose = types.BoolNull()

							if t := ccv.Get("nextHopLoose.optionType"); t.Exists() {
								va := ccv.Get("nextHopLoose.value")
								if t.String() == "global" {
									ccItem.NextHopLoose = types.BoolValue(va.Bool())
								}
							}
							ccItem.Vpn = types.Int64Null()

							if t := ccv.Get("vpn.optionType"); t.Exists() {
								va := ccv.Get("vpn.value")
								if t.String() == "global" {
									ccItem.Vpn = types.Int64Value(va.Int())
								}
							}
							cItem.SetParameters = append(cItem.SetParameters, ccItem)
							return true
						})
					}
					cItem.RedirectDnsField = types.StringNull()

					if t := cv.Get("redirectDns.field.optionType"); t.Exists() {
						va := cv.Get("redirectDns.field.value")
						if t.String() == "global" {
							cItem.RedirectDnsField = types.StringValue(va.String())
						}
					}
					cItem.RedirectDnsValue = types.StringNull()

					if t := cv.Get("redirectDns.value.optionType"); t.Exists() {
						va := cv.Get("redirectDns.value.value")
						if t.String() == "global" {
							cItem.RedirectDnsValue = types.StringValue(va.String())
						}
					}
					cItem.AppqoeTcpOptimization = types.BoolNull()

					if t := cv.Get("appqoeOptimization.tcpOptimization.optionType"); t.Exists() {
						va := cv.Get("appqoeOptimization.tcpOptimization.value")
						if t.String() == "global" {
							cItem.AppqoeTcpOptimization = types.BoolValue(va.Bool())
						}
					}
					cItem.AppqoeDreOptimization = types.BoolNull()

					if t := cv.Get("appqoeOptimization.dreOptimization.optionType"); t.Exists() {
						va := cv.Get("appqoeOptimization.dreOptimization.value")
						if t.String() == "global" {
							cItem.AppqoeDreOptimization = types.BoolValue(va.Bool())
						}
					}
					cItem.AppqoeServiceNodeGroup = types.StringNull()

					if t := cv.Get("appqoeOptimization.serviceNodeGroup.optionType"); t.Exists() {
						va := cv.Get("appqoeOptimization.serviceNodeGroup.value")
						if t.String() == "global" {
							cItem.AppqoeServiceNodeGroup = types.StringValue(va.String())
						}
					}
					cItem.LossCorrectType = types.StringNull()

					if t := cv.Get("lossCorrection.lossCorrectionType.optionType"); t.Exists() {
						va := cv.Get("lossCorrection.lossCorrectionType.value")
						if t.String() == "global" {
							cItem.LossCorrectType = types.StringValue(va.String())
						}
					}
					cItem.LossCorrectFecThreshold = types.Int64Null()

					if t := cv.Get("lossCorrection.lossCorrectFec.optionType"); t.Exists() {
						va := cv.Get("lossCorrection.lossCorrectFec.value")
						if t.String() == "global" {
							cItem.LossCorrectFecThreshold = types.Int64Value(va.Int())
						}
					}
					cItem.Count = types.StringNull()

					if t := cv.Get("count.optionType"); t.Exists() {
						va := cv.Get("count.value")
						if t.String() == "global" {
							cItem.Count = types.StringValue(va.String())
						}
					}
					cItem.Log = types.BoolNull()

					if t := cv.Get("log.optionType"); t.Exists() {
						va := cv.Get("log.value")
						if t.String() == "global" {
							cItem.Log = types.BoolValue(va.Bool())
						}
					}
					cItem.CloudSaas = types.BoolNull()

					if t := cv.Get("cloudSaas.optionType"); t.Exists() {
						va := cv.Get("cloudSaas.value")
						if t.String() == "global" {
							cItem.CloudSaas = types.BoolValue(va.Bool())
						}
					}
					cItem.CloudProbe = types.BoolNull()

					if t := cv.Get("cloudProbe.optionType"); t.Exists() {
						va := cv.Get("cloudProbe.value")
						if t.String() == "global" {
							cItem.CloudProbe = types.BoolValue(va.Bool())
						}
					}
					cItem.Cflowd = types.BoolNull()

					if t := cv.Get("cflowd.optionType"); t.Exists() {
						va := cv.Get("cflowd.value")
						if t.String() == "global" {
							cItem.Cflowd = types.BoolValue(va.Bool())
						}
					}
					cItem.NatPool = types.Int64Null()

					if t := cv.Get("natPool.optionType"); t.Exists() {
						va := cv.Get("natPool.value")
						if t.String() == "global" {
							cItem.NatPool = types.Int64Value(va.Int())
						}
					}
					cItem.NatVpn = types.BoolNull()

					if t := cv.Get("nat.useVpn.optionType"); t.Exists() {
						va := cv.Get("nat.useVpn.value")
						if t.String() == "global" {
							cItem.NatVpn = types.BoolValue(va.Bool())
						}
					}
					cItem.NatFallback = types.BoolNull()

					if t := cv.Get("nat.fallback.optionType"); t.Exists() {
						va := cv.Get("nat.fallback.value")
						if t.String() == "global" {
							cItem.NatFallback = types.BoolValue(va.Bool())
						}
					}
					cItem.NatBypass = types.BoolNull()

					if t := cv.Get("nat.bypass.optionType"); t.Exists() {
						va := cv.Get("nat.bypass.value")
						if t.String() == "global" {
							cItem.NatBypass = types.BoolValue(va.Bool())
						}
					}
					cItem.NatDiaPools = types.SetNull(types.Int64Type)

					if t := cv.Get("nat.diaPool.optionType"); t.Exists() {
						va := cv.Get("nat.diaPool.value")
						if t.String() == "global" {
							cItem.NatDiaPools = helpers.GetInt64Set(va.Array())
						}
					}
					cItem.NatDiaInterfaces = types.SetNull(types.StringType)

					if t := cv.Get("nat.diaInterface.optionType"); t.Exists() {
						va := cv.Get("nat.diaInterface.value")
						if t.String() == "global" {
							cItem.NatDiaInterfaces = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SecureInternetGateway = types.BoolNull()

					if t := cv.Get("sig.optionType"); t.Exists() {
						va := cv.Get("sig.value")
						if t.String() == "global" {
							cItem.SecureInternetGateway = types.BoolValue(va.Bool())
						}
					}
					cItem.FallbackToRouting = types.BoolNull()

					if t := cv.Get("fallbackToRouting.optionType"); t.Exists() {
						va := cv.Get("fallbackToRouting.value")
						if t.String() == "global" {
							cItem.FallbackToRouting = types.BoolValue(va.Bool())
						}
					}
					cItem.SecureServiceEdge = types.BoolNull()

					if t := cv.Get("sse.secureServiceEdge.optionType"); t.Exists() {
						va := cv.Get("sse.secureServiceEdge.value")
						if t.String() == "global" {
							cItem.SecureServiceEdge = types.BoolValue(va.Bool())
						}
					}
					cItem.SecureServiceEdgeInstance = types.StringNull()

					if t := cv.Get("sse.secureServiceEdgeInstance.optionType"); t.Exists() {
						va := cv.Get("sse.secureServiceEdgeInstance.value")
						if t.String() == "global" {
							cItem.SecureServiceEdgeInstance = types.StringValue(va.String())
						}
					}
					item.Actions = append(item.Actions, cItem)
					return true
				})
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	} else {
		data.Sequences = nil
	}
	if !fullRead {
		resultSequences := make([]ApplicationPriorityTrafficPolicySequences, 0, len(data.Sequences))
		matchedSequences := make([]bool, len(data.Sequences))
		for _, oldItem := range oldSequences {
			for ni := range data.Sequences {
				if matchedSequences[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.SequenceId.ValueInt64() != data.Sequences[ni].SequenceId.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedSequences[ni] = true
					{
						resultC := make([]ApplicationPriorityTrafficPolicySequencesMatchEntries, 0, len(data.Sequences[ni].MatchEntries))
						matchedC := make([]bool, len(data.Sequences[ni].MatchEntries))
						for _, oldCItem := range oldItem.MatchEntries {
							for nci := range data.Sequences[ni].MatchEntries {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.ApplicationListId.ValueString() != data.Sequences[ni].MatchEntries[nci].ApplicationListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.SaasApplicationListId.ValueString() != data.Sequences[ni].MatchEntries[nci].SaasApplicationListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.ServiceAreas).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].ServiceAreas).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.TrafficCategory.ValueString() != data.Sequences[ni].MatchEntries[nci].TrafficCategory.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.DnsApplicationListId.ValueString() != data.Sequences[ni].MatchEntries[nci].DnsApplicationListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.TrafficClass.ValueString() != data.Sequences[ni].MatchEntries[nci].TrafficClass.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.Dscps).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].Dscps).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.PacketLength.ValueString() != data.Sequences[ni].MatchEntries[nci].PacketLength.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.Protocols).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].Protocols).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.IcmpMessages).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].IcmpMessages).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.Icmp6Messages).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].Icmp6Messages).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.SourceDataIpv4PrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].SourceDataIpv4PrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.SourceDataIpv6PrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].SourceDataIpv6PrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.SourceIpv4Prefix.ValueString() != data.Sequences[ni].MatchEntries[nci].SourceIpv4Prefix.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.SourceIpv6Prefix.ValueString() != data.Sequences[ni].MatchEntries[nci].SourceIpv6Prefix.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourcePorts).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourcePorts).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.DestinationDataIpv4PrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].DestinationDataIpv4PrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.DestinationDataIpv6PrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].DestinationDataIpv6PrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.DestinationIpv4Prefix.ValueString() != data.Sequences[ni].MatchEntries[nci].DestinationIpv4Prefix.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.DestinationIpv6Prefix.ValueString() != data.Sequences[ni].MatchEntries[nci].DestinationIpv6Prefix.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.DestinationPorts).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].DestinationPorts).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Tcp.ValueString() != data.Sequences[ni].MatchEntries[nci].Tcp.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.DestinationRegion.ValueString() != data.Sequences[ni].MatchEntries[nci].DestinationRegion.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.TrafficTo.ValueString() != data.Sequences[ni].MatchEntries[nci].TrafficTo.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Dns.ValueString() != data.Sequences[ni].MatchEntries[nci].Dns.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.Sequences[ni].MatchEntries[nci])
									break
								}
							}
						}
						for nci := range data.Sequences[ni].MatchEntries {
							if !matchedC[nci] {
								resultC = append(resultC, data.Sequences[ni].MatchEntries[nci])
							}
						}
						data.Sequences[ni].MatchEntries = resultC
					}
					{
						resultC := make([]ApplicationPriorityTrafficPolicySequencesActions, 0, len(data.Sequences[ni].Actions))
						matchedC := make([]bool, len(data.Sequences[ni].Actions))
						for _, oldCItem := range oldItem.Actions {
							for nci := range data.Sequences[ni].Actions {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.BackupSlaPreferredColors).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Actions[nci].BackupSlaPreferredColors).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.RedirectDnsField.ValueString() != data.Sequences[ni].Actions[nci].RedirectDnsField.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.RedirectDnsValue.ValueString() != data.Sequences[ni].Actions[nci].RedirectDnsValue.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.AppqoeTcpOptimization.ValueBool() != data.Sequences[ni].Actions[nci].AppqoeTcpOptimization.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.AppqoeDreOptimization.ValueBool() != data.Sequences[ni].Actions[nci].AppqoeDreOptimization.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.AppqoeServiceNodeGroup.ValueString() != data.Sequences[ni].Actions[nci].AppqoeServiceNodeGroup.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.LossCorrectType.ValueString() != data.Sequences[ni].Actions[nci].LossCorrectType.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.LossCorrectFecThreshold.ValueInt64() != data.Sequences[ni].Actions[nci].LossCorrectFecThreshold.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Count.ValueString() != data.Sequences[ni].Actions[nci].Count.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Log.ValueBool() != data.Sequences[ni].Actions[nci].Log.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.CloudSaas.ValueBool() != data.Sequences[ni].Actions[nci].CloudSaas.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.CloudProbe.ValueBool() != data.Sequences[ni].Actions[nci].CloudProbe.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Cflowd.ValueBool() != data.Sequences[ni].Actions[nci].Cflowd.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.NatPool.ValueInt64() != data.Sequences[ni].Actions[nci].NatPool.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.NatVpn.ValueBool() != data.Sequences[ni].Actions[nci].NatVpn.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.NatFallback.ValueBool() != data.Sequences[ni].Actions[nci].NatFallback.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.NatBypass.ValueBool() != data.Sequences[ni].Actions[nci].NatBypass.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.NatDiaPools).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Actions[nci].NatDiaPools).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.NatDiaInterfaces).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Actions[nci].NatDiaInterfaces).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.SecureInternetGateway.ValueBool() != data.Sequences[ni].Actions[nci].SecureInternetGateway.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.FallbackToRouting.ValueBool() != data.Sequences[ni].Actions[nci].FallbackToRouting.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.SecureServiceEdge.ValueBool() != data.Sequences[ni].Actions[nci].SecureServiceEdge.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.SecureServiceEdgeInstance.ValueString() != data.Sequences[ni].Actions[nci].SecureServiceEdgeInstance.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									{
										resultCC := make([]ApplicationPriorityTrafficPolicySequencesActionsSlaClasses, 0, len(data.Sequences[ni].Actions[nci].SlaClasses))
										matchedCC := make([]bool, len(data.Sequences[ni].Actions[nci].SlaClasses))
										for _, oldCCItem := range oldCItem.SlaClasses {
											for ncci := range data.Sequences[ni].Actions[nci].SlaClasses {
												if matchedCC[ncci] {
													continue
												}
												keyMatchCC := true
												if keyMatchCC {
													if oldCCItem.SlaClassListId.ValueString() != data.Sequences[ni].Actions[nci].SlaClasses[ncci].SlaClassListId.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if helpers.GetStringFromSet(oldCCItem.PreferredColors).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Actions[nci].SlaClasses[ncci].PreferredColors).ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.PreferredColorGroupListId.ValueString() != data.Sequences[ni].Actions[nci].SlaClasses[ncci].PreferredColorGroupListId.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.Strict.ValueBool() != data.Sequences[ni].Actions[nci].SlaClasses[ncci].Strict.ValueBool() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.FallbackToBestPath.ValueBool() != data.Sequences[ni].Actions[nci].SlaClasses[ncci].FallbackToBestPath.ValueBool() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if helpers.GetStringFromSet(oldCCItem.PreferredRemoteColors).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Actions[nci].SlaClasses[ncci].PreferredRemoteColors).ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.RemoteColorRestrict.ValueBool() != data.Sequences[ni].Actions[nci].SlaClasses[ncci].RemoteColorRestrict.ValueBool() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													matchedCC[ncci] = true
													resultCC = append(resultCC, data.Sequences[ni].Actions[nci].SlaClasses[ncci])
													break
												}
											}
										}
										for ncci := range data.Sequences[ni].Actions[nci].SlaClasses {
											if !matchedCC[ncci] {
												resultCC = append(resultCC, data.Sequences[ni].Actions[nci].SlaClasses[ncci])
											}
										}
										data.Sequences[ni].Actions[nci].SlaClasses = resultCC
									}
									{
										resultCC := make([]ApplicationPriorityTrafficPolicySequencesActionsSetParameters, 0, len(data.Sequences[ni].Actions[nci].SetParameters))
										matchedCC := make([]bool, len(data.Sequences[ni].Actions[nci].SetParameters))
										for _, oldCCItem := range oldCItem.SetParameters {
											for ncci := range data.Sequences[ni].Actions[nci].SetParameters {
												if matchedCC[ncci] {
													continue
												}
												keyMatchCC := true
												if keyMatchCC {
													if oldCCItem.Dscp.ValueInt64() != data.Sequences[ni].Actions[nci].SetParameters[ncci].Dscp.ValueInt64() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.PolicerId.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].PolicerId.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.PreferredColorGroupId.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].PreferredColorGroupId.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ForwardingClassListId.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ForwardingClassListId.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if helpers.GetStringFromSet(oldCCItem.LocalTlocListColors).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Actions[nci].SetParameters[ncci].LocalTlocListColors).ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.LocalTlocListRestrict.ValueBool() != data.Sequences[ni].Actions[nci].SetParameters[ncci].LocalTlocListRestrict.ValueBool() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.LocalTlocListEncapsulation.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].LocalTlocListEncapsulation.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if helpers.GetStringFromSet(oldCCItem.PreferredRemoteColors).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Actions[nci].SetParameters[ncci].PreferredRemoteColors).ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.PreferredRemoteColorRestrict.ValueBool() != data.Sequences[ni].Actions[nci].SetParameters[ncci].PreferredRemoteColorRestrict.ValueBool() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if helpers.GetStringFromSet(oldCCItem.TlocColor).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Actions[nci].SetParameters[ncci].TlocColor).ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.TlocEncapsulation.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].TlocEncapsulation.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.TlocIp.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].TlocIp.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.TlocListId.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].TlocListId.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if helpers.GetStringFromSet(oldCCItem.ServiceTlocColor).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceTlocColor).ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceTlocEncapsulation.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceTlocEncapsulation.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceTlocIp.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceTlocIp.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceVpn.ValueInt64() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceVpn.ValueInt64() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceType.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceType.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceLocal.ValueBool() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceLocal.ValueBool() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceRestrict.ValueBool() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceRestrict.ValueBool() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceTlocListId.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceTlocListId.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainType.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceChainType.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainVpn.ValueInt64() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceChainVpn.ValueInt64() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainLocal.ValueBool() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceChainLocal.ValueBool() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainFallbackToRouting.ValueBool() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceChainFallbackToRouting.ValueBool() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if helpers.GetStringFromSet(oldCCItem.ServiceChainTlocColor).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceChainTlocColor).ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainTlocEncapsulation.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceChainTlocEncapsulation.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainTlocIp.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceChainTlocIp.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainTlocListId.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].ServiceChainTlocListId.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.NextHopIpv4.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].NextHopIpv4.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.NextHopIpv6.ValueString() != data.Sequences[ni].Actions[nci].SetParameters[ncci].NextHopIpv6.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.NextHopLoose.ValueBool() != data.Sequences[ni].Actions[nci].SetParameters[ncci].NextHopLoose.ValueBool() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.Vpn.ValueInt64() != data.Sequences[ni].Actions[nci].SetParameters[ncci].Vpn.ValueInt64() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													matchedCC[ncci] = true
													resultCC = append(resultCC, data.Sequences[ni].Actions[nci].SetParameters[ncci])
													break
												}
											}
										}
										for ncci := range data.Sequences[ni].Actions[nci].SetParameters {
											if !matchedCC[ncci] {
												resultCC = append(resultCC, data.Sequences[ni].Actions[nci].SetParameters[ncci])
											}
										}
										data.Sequences[ni].Actions[nci].SetParameters = resultCC
									}
									resultC = append(resultC, data.Sequences[ni].Actions[nci])
									break
								}
							}
						}
						for nci := range data.Sequences[ni].Actions {
							if !matchedC[nci] {
								resultC = append(resultC, data.Sequences[ni].Actions[nci])
							}
						}
						data.Sequences[ni].Actions = resultC
					}
					resultSequences = append(resultSequences, data.Sequences[ni])
					break
				}
			}
		}
		for ni := range data.Sequences {
			if !matchedSequences[ni] {
				resultSequences = append(resultSequences, data.Sequences[ni])
			}
		}
		data.Sequences = resultSequences
	}
}

// End of section. //template:end fromBody
