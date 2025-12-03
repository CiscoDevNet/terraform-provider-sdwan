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
	ServiceArea                     types.Set    `tfsdk:"service_area"`
	TrafficCategory                 types.String `tfsdk:"traffic_category"`
	DnsApplicationListId            types.String `tfsdk:"dns_application_list_id"`
	TrafficClass                    types.String `tfsdk:"traffic_class"`
	Dscp                            types.Int64  `tfsdk:"dscp"`
	PacketLength                    types.String `tfsdk:"packet_length"`
	Protocols                       types.Set    `tfsdk:"protocols"`
	IcmpMessage                     types.Set    `tfsdk:"icmp_message"`
	Icmp6Message                    types.Set    `tfsdk:"icmp6_message"`
	SourceDataIpv4PrefxListId       types.String `tfsdk:"source_data_ipv4_prefx_list_id"`
	SourceDataIpv6PrefxListId       types.String `tfsdk:"source_data_ipv6_prefx_list_id"`
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
	SlaClasses              []ApplicationPriorityTrafficPolicySequencesActionsSlaClasses    `tfsdk:"sla_classes"`
	BackupSlaPreferredColor types.Set                                                       `tfsdk:"backup_sla_preferred_color"`
	SetParameters           []ApplicationPriorityTrafficPolicySequencesActionsSetParameters `tfsdk:"set_parameters"`
	RedirectDnsField        types.String                                                    `tfsdk:"redirect_dns_field"`
	RedirectDnsValue        types.String                                                    `tfsdk:"redirect_dns_value"`
	LossCorrectType         types.String                                                    `tfsdk:"loss_correct_type"`
	LossCorrectFecThreshold types.Int64                                                     `tfsdk:"loss_correct_fec_threshold"`
	Count                   types.String                                                    `tfsdk:"count"`
	Log                     types.Bool                                                      `tfsdk:"log"`
	CloudSaas               types.Bool                                                      `tfsdk:"cloud_saas"`
	CloudProbe              types.Bool                                                      `tfsdk:"cloud_probe"`
	NatPool                 types.Int64                                                     `tfsdk:"nat_pool"`
	NatVpn                  types.Bool                                                      `tfsdk:"nat_vpn"`
	NatFallback             types.Bool                                                      `tfsdk:"nat_fallback"`
	NatBypass               types.Bool                                                      `tfsdk:"nat_bypass"`
	NatDiaPool              types.Set                                                       `tfsdk:"nat_dia_pool"`
	NatDiaInterface         types.Set                                                       `tfsdk:"nat_dia_interface"`
	SecureInternetGateway   types.Bool                                                      `tfsdk:"secure_internet_gateway"`
	FallbackToRouting       types.Bool                                                      `tfsdk:"fallback_to_routing"`
}

type ApplicationPriorityTrafficPolicySequencesActionsSlaClasses struct {
	SlaClassListId            types.String `tfsdk:"sla_class_list_id"`
	PreferredColor            types.Set    `tfsdk:"preferred_color"`
	PreferredColorGroupListId types.String `tfsdk:"preferred_color_group_list_id"`
	Strict                    types.Bool   `tfsdk:"strict"`
	FallbackToBestPath        types.Bool   `tfsdk:"fallback_to_best_path"`
	PreferredRemoteColor      types.Set    `tfsdk:"preferred_remote_color"`
	RemoteColorRestrict       types.Bool   `tfsdk:"remote_color_restrict"`
}
type ApplicationPriorityTrafficPolicySequencesActionsSetParameters struct {
	Dscp                          types.Int64  `tfsdk:"dscp"`
	PolicerId                     types.String `tfsdk:"policer_id"`
	PreferredColorGroupId         types.String `tfsdk:"preferred_color_group_id"`
	ForwardingClassListId         types.String `tfsdk:"forwarding_class_list_id"`
	LocalTlocListColor            types.Set    `tfsdk:"local_tloc_list_color"`
	LocalTlocListRestrict         types.Bool   `tfsdk:"local_tloc_list_restrict"`
	LocalTlocListEncapsulation    types.String `tfsdk:"local_tloc_list_encapsulation"`
	PreferredRemoteColorId        types.Set    `tfsdk:"preferred_remote_color_id"`
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
					if !childItem.ServiceArea.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "serviceArea.optionType", "global")
							var values []string
							childItem.ServiceArea.ElementsAs(ctx, &values, false)
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
					if !childItem.Dscp.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "dscp.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "dscp.value", childItem.Dscp.ValueInt64())
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
					if !childItem.IcmpMessage.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "icmpMessage.optionType", "global")
							var values []string
							childItem.IcmpMessage.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "icmpMessage.value", values)
						}
					}
					if !childItem.Icmp6Message.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "icmp6Message.optionType", "global")
							var values []string
							childItem.Icmp6Message.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "icmp6Message.value", values)
						}
					}
					if !childItem.SourceDataIpv4PrefxListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefixList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefixList.refId.value", childItem.SourceDataIpv4PrefxListId.ValueString())
						}
					}
					if !childItem.SourceDataIpv6PrefxListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataIpv6PrefixList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataIpv6PrefixList.refId.value", childItem.SourceDataIpv6PrefxListId.ValueString())
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
							if !childChildItem.PreferredColor.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredColor.optionType", "global")
									var values []string
									childChildItem.PreferredColor.ElementsAs(ctx, &values, false)
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
							if !childChildItem.PreferredRemoteColor.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredRemoteColor.optionType", "global")
									var values []string
									childChildItem.PreferredRemoteColor.ElementsAs(ctx, &values, false)
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
					if !childItem.BackupSlaPreferredColor.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "backupSlaPreferredColor.optionType", "global")
							var values []string
							childItem.BackupSlaPreferredColor.ElementsAs(ctx, &values, false)
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
							if !childChildItem.LocalTlocListColor.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.color.optionType", "global")
									var values []string
									childChildItem.LocalTlocListColor.ElementsAs(ctx, &values, false)
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
							if !childChildItem.PreferredRemoteColorId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredRemoteColor.color.optionType", "global")
									var values []string
									childChildItem.PreferredRemoteColorId.ElementsAs(ctx, &values, false)
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
					if !childItem.NatDiaPool.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.diaPool.optionType", "global")
							var values []int64
							childItem.NatDiaPool.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.diaPool.value", values)
						}
					}
					if !childItem.NatDiaInterface.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "nat.diaInterface.optionType", "global")
							var values []string
							childItem.NatDiaInterface.ElementsAs(ctx, &values, false)
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
func (data *ApplicationPriorityTrafficPolicy) fromBody(ctx context.Context, res gjson.Result) {
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
					cItem.ServiceArea = types.SetNull(types.StringType)

					if t := cv.Get("serviceArea.optionType"); t.Exists() {
						va := cv.Get("serviceArea.value")
						if t.String() == "global" {
							cItem.ServiceArea = helpers.GetStringSet(va.Array())
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
					cItem.Dscp = types.Int64Null()

					if t := cv.Get("dscp.optionType"); t.Exists() {
						va := cv.Get("dscp.value")
						if t.String() == "global" {
							cItem.Dscp = types.Int64Value(va.Int())
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
					cItem.IcmpMessage = types.SetNull(types.StringType)

					if t := cv.Get("icmpMessage.optionType"); t.Exists() {
						va := cv.Get("icmpMessage.value")
						if t.String() == "global" {
							cItem.IcmpMessage = helpers.GetStringSet(va.Array())
						}
					}
					cItem.Icmp6Message = types.SetNull(types.StringType)

					if t := cv.Get("icmp6Message.optionType"); t.Exists() {
						va := cv.Get("icmp6Message.value")
						if t.String() == "global" {
							cItem.Icmp6Message = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceDataIpv4PrefxListId = types.StringNull()

					if t := cv.Get("sourceDataPrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("sourceDataPrefixList.refId.value")
						if t.String() == "global" {
							cItem.SourceDataIpv4PrefxListId = types.StringValue(va.String())
						}
					}
					cItem.SourceDataIpv6PrefxListId = types.StringNull()

					if t := cv.Get("sourceDataIpv6PrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("sourceDataIpv6PrefixList.refId.value")
						if t.String() == "global" {
							cItem.SourceDataIpv6PrefxListId = types.StringValue(va.String())
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
							ccItem.PreferredColor = types.SetNull(types.StringType)

							if t := ccv.Get("preferredColor.optionType"); t.Exists() {
								va := ccv.Get("preferredColor.value")
								if t.String() == "global" {
									ccItem.PreferredColor = helpers.GetStringSet(va.Array())
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
							ccItem.PreferredRemoteColor = types.SetNull(types.StringType)

							if t := ccv.Get("preferredRemoteColor.optionType"); t.Exists() {
								va := ccv.Get("preferredRemoteColor.value")
								if t.String() == "global" {
									ccItem.PreferredRemoteColor = helpers.GetStringSet(va.Array())
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
					cItem.BackupSlaPreferredColor = types.SetNull(types.StringType)

					if t := cv.Get("backupSlaPreferredColor.optionType"); t.Exists() {
						va := cv.Get("backupSlaPreferredColor.value")
						if t.String() == "global" {
							cItem.BackupSlaPreferredColor = helpers.GetStringSet(va.Array())
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
							ccItem.LocalTlocListColor = types.SetNull(types.StringType)

							if t := ccv.Get("localTlocList.color.optionType"); t.Exists() {
								va := ccv.Get("localTlocList.color.value")
								if t.String() == "global" {
									ccItem.LocalTlocListColor = helpers.GetStringSet(va.Array())
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
							ccItem.PreferredRemoteColorId = types.SetNull(types.StringType)

							if t := ccv.Get("preferredRemoteColor.color.optionType"); t.Exists() {
								va := ccv.Get("preferredRemoteColor.color.value")
								if t.String() == "global" {
									ccItem.PreferredRemoteColorId = helpers.GetStringSet(va.Array())
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
					cItem.NatDiaPool = types.SetNull(types.Int64Type)

					if t := cv.Get("nat.diaPool.optionType"); t.Exists() {
						va := cv.Get("nat.diaPool.value")
						if t.String() == "global" {
							cItem.NatDiaPool = helpers.GetInt64Set(va.Array())
						}
					}
					cItem.NatDiaInterface = types.SetNull(types.StringType)

					if t := cv.Get("nat.diaInterface.optionType"); t.Exists() {
						va := cv.Get("nat.diaInterface.value")
						if t.String() == "global" {
							cItem.NatDiaInterface = helpers.GetStringSet(va.Array())
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
					item.Actions = append(item.Actions, cItem)
					return true
				})
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ApplicationPriorityTrafficPolicy) updateFromBody(ctx context.Context, res gjson.Result) {
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
	for i := range data.Sequences {
		keys := [...]string{"sequenceId"}
		keyValues := [...]string{strconv.FormatInt(data.Sequences[i].SequenceId.ValueInt64(), 10)}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "sequences").ForEach(
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
		data.Sequences[i].SequenceId = types.Int64Null()

		if t := r.Get("sequenceId.optionType"); t.Exists() {
			va := r.Get("sequenceId.value")
			if t.String() == "global" {
				data.Sequences[i].SequenceId = types.Int64Value(va.Int())
			}
		}
		data.Sequences[i].SequenceName = types.StringNull()

		if t := r.Get("sequenceName.optionType"); t.Exists() {
			va := r.Get("sequenceName.value")
			if t.String() == "global" {
				data.Sequences[i].SequenceName = types.StringValue(va.String())
			}
		}
		data.Sequences[i].BaseAction = types.StringNull()

		if t := r.Get("baseAction.optionType"); t.Exists() {
			va := r.Get("baseAction.value")
			if t.String() == "global" {
				data.Sequences[i].BaseAction = types.StringValue(va.String())
			}
		}
		data.Sequences[i].Protocol = types.StringNull()

		if t := r.Get("sequenceIpType.optionType"); t.Exists() {
			va := r.Get("sequenceIpType.value")
			if t.String() == "global" {
				data.Sequences[i].Protocol = types.StringValue(va.String())
			}
		}
		for ci := range data.Sequences[i].MatchEntries {
			keys := [...]string{"appList.refId", "saasAppList.refId", "serviceArea", "trafficCategory", "dnsAppList.refId", "trafficClass", "dscp", "packetLength", "protocol", "icmpMessage", "icmp6Message", "sourceDataPrefixList.refId", "sourceDataIpv6PrefixList.refId", "sourceIp", "sourceIpv6", "sourcePort", "destinationDataPrefixList.refId", "destinationDataIpv6PrefixList.refId", "destinationIp", "destinationIpv6", "destinationPort", "tcp", "destinationRegion", "trafficTo", "dns"}
			keyValues := [...]string{data.Sequences[i].MatchEntries[ci].ApplicationListId.ValueString(), data.Sequences[i].MatchEntries[ci].SaasApplicationListId.ValueString(), helpers.GetStringFromSet(data.Sequences[i].MatchEntries[ci].ServiceArea).ValueString(), data.Sequences[i].MatchEntries[ci].TrafficCategory.ValueString(), data.Sequences[i].MatchEntries[ci].DnsApplicationListId.ValueString(), data.Sequences[i].MatchEntries[ci].TrafficClass.ValueString(), strconv.FormatInt(data.Sequences[i].MatchEntries[ci].Dscp.ValueInt64(), 10), data.Sequences[i].MatchEntries[ci].PacketLength.ValueString(), helpers.GetStringFromSet(data.Sequences[i].MatchEntries[ci].Protocols).ValueString(), helpers.GetStringFromSet(data.Sequences[i].MatchEntries[ci].IcmpMessage).ValueString(), helpers.GetStringFromSet(data.Sequences[i].MatchEntries[ci].Icmp6Message).ValueString(), data.Sequences[i].MatchEntries[ci].SourceDataIpv4PrefxListId.ValueString(), data.Sequences[i].MatchEntries[ci].SourceDataIpv6PrefxListId.ValueString(), data.Sequences[i].MatchEntries[ci].SourceIpv4Prefix.ValueString(), data.Sequences[i].MatchEntries[ci].SourceIpv6Prefix.ValueString(), helpers.GetStringFromSet(data.Sequences[i].MatchEntries[ci].SourcePorts).ValueString(), data.Sequences[i].MatchEntries[ci].DestinationDataIpv4PrefixListId.ValueString(), data.Sequences[i].MatchEntries[ci].DestinationDataIpv6PrefixListId.ValueString(), data.Sequences[i].MatchEntries[ci].DestinationIpv4Prefix.ValueString(), data.Sequences[i].MatchEntries[ci].DestinationIpv6Prefix.ValueString(), helpers.GetStringFromSet(data.Sequences[i].MatchEntries[ci].DestinationPorts).ValueString(), data.Sequences[i].MatchEntries[ci].Tcp.ValueString(), data.Sequences[i].MatchEntries[ci].DestinationRegion.ValueString(), data.Sequences[i].MatchEntries[ci].TrafficTo.ValueString(), data.Sequences[i].MatchEntries[ci].Dns.ValueString()}
			keyValuesVariables := [...]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}

			var cr gjson.Result
			r.Get("match.entries").ForEach(
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
			data.Sequences[i].MatchEntries[ci].ApplicationListId = types.StringNull()

			if t := cr.Get("appList.refId.optionType"); t.Exists() {
				va := cr.Get("appList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].ApplicationListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].SaasApplicationListId = types.StringNull()

			if t := cr.Get("saasAppList.refId.optionType"); t.Exists() {
				va := cr.Get("saasAppList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].SaasApplicationListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].ServiceArea = types.SetNull(types.StringType)

			if t := cr.Get("serviceArea.optionType"); t.Exists() {
				va := cr.Get("serviceArea.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].ServiceArea = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].MatchEntries[ci].TrafficCategory = types.StringNull()

			if t := cr.Get("trafficCategory.optionType"); t.Exists() {
				va := cr.Get("trafficCategory.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].TrafficCategory = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].DnsApplicationListId = types.StringNull()

			if t := cr.Get("dnsAppList.refId.optionType"); t.Exists() {
				va := cr.Get("dnsAppList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].DnsApplicationListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].TrafficClass = types.StringNull()

			if t := cr.Get("trafficClass.optionType"); t.Exists() {
				va := cr.Get("trafficClass.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].TrafficClass = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].Dscp = types.Int64Null()

			if t := cr.Get("dscp.optionType"); t.Exists() {
				va := cr.Get("dscp.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Dscp = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].MatchEntries[ci].PacketLength = types.StringNull()

			if t := cr.Get("packetLength.optionType"); t.Exists() {
				va := cr.Get("packetLength.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].PacketLength = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].Protocols = types.SetNull(types.StringType)

			if t := cr.Get("protocol.optionType"); t.Exists() {
				va := cr.Get("protocol.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Protocols = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].MatchEntries[ci].IcmpMessage = types.SetNull(types.StringType)

			if t := cr.Get("icmpMessage.optionType"); t.Exists() {
				va := cr.Get("icmpMessage.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].IcmpMessage = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].MatchEntries[ci].Icmp6Message = types.SetNull(types.StringType)

			if t := cr.Get("icmp6Message.optionType"); t.Exists() {
				va := cr.Get("icmp6Message.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Icmp6Message = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].MatchEntries[ci].SourceDataIpv4PrefxListId = types.StringNull()

			if t := cr.Get("sourceDataPrefixList.refId.optionType"); t.Exists() {
				va := cr.Get("sourceDataPrefixList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].SourceDataIpv4PrefxListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].SourceDataIpv6PrefxListId = types.StringNull()

			if t := cr.Get("sourceDataIpv6PrefixList.refId.optionType"); t.Exists() {
				va := cr.Get("sourceDataIpv6PrefixList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].SourceDataIpv6PrefxListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].SourceIpv4Prefix = types.StringNull()

			if t := cr.Get("sourceIp.optionType"); t.Exists() {
				va := cr.Get("sourceIp.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].SourceIpv4Prefix = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].SourceIpv6Prefix = types.StringNull()

			if t := cr.Get("sourceIpv6.optionType"); t.Exists() {
				va := cr.Get("sourceIpv6.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].SourceIpv6Prefix = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].SourcePorts = types.SetNull(types.StringType)

			if t := cr.Get("sourcePort.optionType"); t.Exists() {
				va := cr.Get("sourcePort.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].SourcePorts = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].MatchEntries[ci].DestinationDataIpv4PrefixListId = types.StringNull()

			if t := cr.Get("destinationDataPrefixList.refId.optionType"); t.Exists() {
				va := cr.Get("destinationDataPrefixList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].DestinationDataIpv4PrefixListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].DestinationDataIpv6PrefixListId = types.StringNull()

			if t := cr.Get("destinationDataIpv6PrefixList.refId.optionType"); t.Exists() {
				va := cr.Get("destinationDataIpv6PrefixList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].DestinationDataIpv6PrefixListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].DestinationIpv4Prefix = types.StringNull()

			if t := cr.Get("destinationIp.optionType"); t.Exists() {
				va := cr.Get("destinationIp.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].DestinationIpv4Prefix = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].DestinationIpv6Prefix = types.StringNull()

			if t := cr.Get("destinationIpv6.optionType"); t.Exists() {
				va := cr.Get("destinationIpv6.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].DestinationIpv6Prefix = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].DestinationPorts = types.SetNull(types.StringType)

			if t := cr.Get("destinationPort.optionType"); t.Exists() {
				va := cr.Get("destinationPort.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].DestinationPorts = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].MatchEntries[ci].Tcp = types.StringNull()

			if t := cr.Get("tcp.optionType"); t.Exists() {
				va := cr.Get("tcp.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Tcp = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].DestinationRegion = types.StringNull()

			if t := cr.Get("destinationRegion.optionType"); t.Exists() {
				va := cr.Get("destinationRegion.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].DestinationRegion = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].TrafficTo = types.StringNull()

			if t := cr.Get("trafficTo.optionType"); t.Exists() {
				va := cr.Get("trafficTo.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].TrafficTo = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].Dns = types.StringNull()

			if t := cr.Get("dns.optionType"); t.Exists() {
				va := cr.Get("dns.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Dns = types.StringValue(va.String())
				}
			}
		}
		for ci := range data.Sequences[i].Actions {
			keys := [...]string{"backupSlaPreferredColor", "redirectDns.field", "redirectDns.value", "lossCorrection.lossCorrectionType", "lossCorrection.lossCorrectFec", "count", "log", "cloudSaas", "cloudProbe", "natPool", "nat.useVpn", "nat.fallback", "nat.bypass", "nat.diaPool", "nat.diaInterface", "sig", "fallbackToRouting"}
			keyValues := [...]string{helpers.GetStringFromSet(data.Sequences[i].Actions[ci].BackupSlaPreferredColor).ValueString(), data.Sequences[i].Actions[ci].RedirectDnsField.ValueString(), data.Sequences[i].Actions[ci].RedirectDnsValue.ValueString(), data.Sequences[i].Actions[ci].LossCorrectType.ValueString(), strconv.FormatInt(data.Sequences[i].Actions[ci].LossCorrectFecThreshold.ValueInt64(), 10), data.Sequences[i].Actions[ci].Count.ValueString(), strconv.FormatBool(data.Sequences[i].Actions[ci].Log.ValueBool()), strconv.FormatBool(data.Sequences[i].Actions[ci].CloudSaas.ValueBool()), strconv.FormatBool(data.Sequences[i].Actions[ci].CloudProbe.ValueBool()), strconv.FormatInt(data.Sequences[i].Actions[ci].NatPool.ValueInt64(), 10), strconv.FormatBool(data.Sequences[i].Actions[ci].NatVpn.ValueBool()), strconv.FormatBool(data.Sequences[i].Actions[ci].NatFallback.ValueBool()), strconv.FormatBool(data.Sequences[i].Actions[ci].NatBypass.ValueBool()), helpers.GetStringFromSet(data.Sequences[i].Actions[ci].NatDiaPool).ValueString(), helpers.GetStringFromSet(data.Sequences[i].Actions[ci].NatDiaInterface).ValueString(), strconv.FormatBool(data.Sequences[i].Actions[ci].SecureInternetGateway.ValueBool()), strconv.FormatBool(data.Sequences[i].Actions[ci].FallbackToRouting.ValueBool())}
			keyValuesVariables := [...]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}

			var cr gjson.Result
			r.Get("actions").ForEach(
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
			for cci := range data.Sequences[i].Actions[ci].SlaClasses {
				keys := [...]string{"slaName.refId", "preferredColor", "preferredColorGroup.refId", "strict", "fallbackToBestPath", "preferredRemoteColor", "remoteColorRestrict"}
				keyValues := [...]string{data.Sequences[i].Actions[ci].SlaClasses[cci].SlaClassListId.ValueString(), helpers.GetStringFromSet(data.Sequences[i].Actions[ci].SlaClasses[cci].PreferredColor).ValueString(), data.Sequences[i].Actions[ci].SlaClasses[cci].PreferredColorGroupListId.ValueString(), strconv.FormatBool(data.Sequences[i].Actions[ci].SlaClasses[cci].Strict.ValueBool()), strconv.FormatBool(data.Sequences[i].Actions[ci].SlaClasses[cci].FallbackToBestPath.ValueBool()), helpers.GetStringFromSet(data.Sequences[i].Actions[ci].SlaClasses[cci].PreferredRemoteColor).ValueString(), strconv.FormatBool(data.Sequences[i].Actions[ci].SlaClasses[cci].RemoteColorRestrict.ValueBool())}
				keyValuesVariables := [...]string{"", "", "", "", "", "", ""}

				var ccr gjson.Result
				cr.Get("slaClass").ForEach(
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
							ccr = v
							return false
						}
						return true
					},
				)
				data.Sequences[i].Actions[ci].SlaClasses[cci].SlaClassListId = types.StringNull()

				if t := ccr.Get("slaName.refId.optionType"); t.Exists() {
					va := ccr.Get("slaName.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SlaClasses[cci].SlaClassListId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SlaClasses[cci].PreferredColor = types.SetNull(types.StringType)

				if t := ccr.Get("preferredColor.optionType"); t.Exists() {
					va := ccr.Get("preferredColor.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SlaClasses[cci].PreferredColor = helpers.GetStringSet(va.Array())
					}
				}
				data.Sequences[i].Actions[ci].SlaClasses[cci].PreferredColorGroupListId = types.StringNull()

				if t := ccr.Get("preferredColorGroup.refId.optionType"); t.Exists() {
					va := ccr.Get("preferredColorGroup.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SlaClasses[cci].PreferredColorGroupListId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SlaClasses[cci].Strict = types.BoolNull()

				if t := ccr.Get("strict.optionType"); t.Exists() {
					va := ccr.Get("strict.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SlaClasses[cci].Strict = types.BoolValue(va.Bool())
					}
				}
				data.Sequences[i].Actions[ci].SlaClasses[cci].FallbackToBestPath = types.BoolNull()

				if t := ccr.Get("fallbackToBestPath.optionType"); t.Exists() {
					va := ccr.Get("fallbackToBestPath.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SlaClasses[cci].FallbackToBestPath = types.BoolValue(va.Bool())
					}
				}
				data.Sequences[i].Actions[ci].SlaClasses[cci].PreferredRemoteColor = types.SetNull(types.StringType)

				if t := ccr.Get("preferredRemoteColor.optionType"); t.Exists() {
					va := ccr.Get("preferredRemoteColor.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SlaClasses[cci].PreferredRemoteColor = helpers.GetStringSet(va.Array())
					}
				}
				data.Sequences[i].Actions[ci].SlaClasses[cci].RemoteColorRestrict = types.BoolNull()

				if t := ccr.Get("remoteColorRestrict.optionType"); t.Exists() {
					va := ccr.Get("remoteColorRestrict.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SlaClasses[cci].RemoteColorRestrict = types.BoolValue(va.Bool())
					}
				}
			}
			data.Sequences[i].Actions[ci].BackupSlaPreferredColor = types.SetNull(types.StringType)

			if t := cr.Get("backupSlaPreferredColor.optionType"); t.Exists() {
				va := cr.Get("backupSlaPreferredColor.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].BackupSlaPreferredColor = helpers.GetStringSet(va.Array())
				}
			}
			for cci := range data.Sequences[i].Actions[ci].SetParameters {
				keys := [...]string{"dscp", "policer.refId", "preferredColorGroup.refId", "forwardingClass.refId", "localTlocList.color", "localTlocList.restrict", "localTlocList.encap", "preferredRemoteColor.color", "preferredRemoteColor.remoteColorRestrict", "tloc.color", "tloc.encap", "tloc.ip", "tlocList.refId", "service.tloc.color", "service.tloc.encap", "service.tloc.ip", "service.vpn", "service.type", "service.local", "service.restrict", "service.tlocList.refId", "serviceChain.type", "serviceChain.vpn", "serviceChain.local", "serviceChain.restrict", "serviceChain.tloc.color", "serviceChain.tloc.encap", "serviceChain.tloc.ip", "serviceChain.tlocList.refId", "nextHop", "nextHopIpv6", "nextHopLoose", "vpn"}
				keyValues := [...]string{strconv.FormatInt(data.Sequences[i].Actions[ci].SetParameters[cci].Dscp.ValueInt64(), 10), data.Sequences[i].Actions[ci].SetParameters[cci].PolicerId.ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].PreferredColorGroupId.ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].ForwardingClassListId.ValueString(), helpers.GetStringFromSet(data.Sequences[i].Actions[ci].SetParameters[cci].LocalTlocListColor).ValueString(), strconv.FormatBool(data.Sequences[i].Actions[ci].SetParameters[cci].LocalTlocListRestrict.ValueBool()), data.Sequences[i].Actions[ci].SetParameters[cci].LocalTlocListEncapsulation.ValueString(), helpers.GetStringFromSet(data.Sequences[i].Actions[ci].SetParameters[cci].PreferredRemoteColorId).ValueString(), strconv.FormatBool(data.Sequences[i].Actions[ci].SetParameters[cci].PreferredRemoteColorRestrict.ValueBool()), helpers.GetStringFromSet(data.Sequences[i].Actions[ci].SetParameters[cci].TlocColor).ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].TlocEncapsulation.ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].TlocIp.ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].TlocListId.ValueString(), helpers.GetStringFromSet(data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocColor).ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocEncapsulation.ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocIp.ValueString(), strconv.FormatInt(data.Sequences[i].Actions[ci].SetParameters[cci].ServiceVpn.ValueInt64(), 10), data.Sequences[i].Actions[ci].SetParameters[cci].ServiceType.ValueString(), strconv.FormatBool(data.Sequences[i].Actions[ci].SetParameters[cci].ServiceLocal.ValueBool()), strconv.FormatBool(data.Sequences[i].Actions[ci].SetParameters[cci].ServiceRestrict.ValueBool()), data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocListId.ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainType.ValueString(), strconv.FormatInt(data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainVpn.ValueInt64(), 10), strconv.FormatBool(data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainLocal.ValueBool()), strconv.FormatBool(data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainFallbackToRouting.ValueBool()), helpers.GetStringFromSet(data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocColor).ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocEncapsulation.ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocIp.ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocListId.ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].NextHopIpv4.ValueString(), data.Sequences[i].Actions[ci].SetParameters[cci].NextHopIpv6.ValueString(), strconv.FormatBool(data.Sequences[i].Actions[ci].SetParameters[cci].NextHopLoose.ValueBool()), strconv.FormatInt(data.Sequences[i].Actions[ci].SetParameters[cci].Vpn.ValueInt64(), 10)}
				keyValuesVariables := [...]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}

				var ccr gjson.Result
				cr.Get("set").ForEach(
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
							ccr = v
							return false
						}
						return true
					},
				)
				data.Sequences[i].Actions[ci].SetParameters[cci].Dscp = types.Int64Null()

				if t := ccr.Get("dscp.optionType"); t.Exists() {
					va := ccr.Get("dscp.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].Dscp = types.Int64Value(va.Int())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].PolicerId = types.StringNull()

				if t := ccr.Get("policer.refId.optionType"); t.Exists() {
					va := ccr.Get("policer.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].PolicerId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].PreferredColorGroupId = types.StringNull()

				if t := ccr.Get("preferredColorGroup.refId.optionType"); t.Exists() {
					va := ccr.Get("preferredColorGroup.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].PreferredColorGroupId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ForwardingClassListId = types.StringNull()

				if t := ccr.Get("forwardingClass.refId.optionType"); t.Exists() {
					va := ccr.Get("forwardingClass.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ForwardingClassListId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].LocalTlocListColor = types.SetNull(types.StringType)

				if t := ccr.Get("localTlocList.color.optionType"); t.Exists() {
					va := ccr.Get("localTlocList.color.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].LocalTlocListColor = helpers.GetStringSet(va.Array())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].LocalTlocListRestrict = types.BoolNull()

				if t := ccr.Get("localTlocList.restrict.optionType"); t.Exists() {
					va := ccr.Get("localTlocList.restrict.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].LocalTlocListRestrict = types.BoolValue(va.Bool())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].LocalTlocListEncapsulation = types.StringNull()

				if t := ccr.Get("localTlocList.encap.optionType"); t.Exists() {
					va := ccr.Get("localTlocList.encap.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].LocalTlocListEncapsulation = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].PreferredRemoteColorId = types.SetNull(types.StringType)

				if t := ccr.Get("preferredRemoteColor.color.optionType"); t.Exists() {
					va := ccr.Get("preferredRemoteColor.color.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].PreferredRemoteColorId = helpers.GetStringSet(va.Array())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].PreferredRemoteColorRestrict = types.BoolNull()

				if t := ccr.Get("preferredRemoteColor.remoteColorRestrict.optionType"); t.Exists() {
					va := ccr.Get("preferredRemoteColor.remoteColorRestrict.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].PreferredRemoteColorRestrict = types.BoolValue(va.Bool())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].TlocColor = types.SetNull(types.StringType)

				if t := ccr.Get("tloc.color.optionType"); t.Exists() {
					va := ccr.Get("tloc.color.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].TlocColor = helpers.GetStringSet(va.Array())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].TlocEncapsulation = types.StringNull()

				if t := ccr.Get("tloc.encap.optionType"); t.Exists() {
					va := ccr.Get("tloc.encap.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].TlocEncapsulation = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].TlocIp = types.StringNull()

				if t := ccr.Get("tloc.ip.optionType"); t.Exists() {
					va := ccr.Get("tloc.ip.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].TlocIp = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].TlocListId = types.StringNull()

				if t := ccr.Get("tlocList.refId.optionType"); t.Exists() {
					va := ccr.Get("tlocList.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].TlocListId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocColor = types.SetNull(types.StringType)

				if t := ccr.Get("service.tloc.color.optionType"); t.Exists() {
					va := ccr.Get("service.tloc.color.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocColor = helpers.GetStringSet(va.Array())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocEncapsulation = types.StringNull()

				if t := ccr.Get("service.tloc.encap.optionType"); t.Exists() {
					va := ccr.Get("service.tloc.encap.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocEncapsulation = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocIp = types.StringNull()

				if t := ccr.Get("service.tloc.ip.optionType"); t.Exists() {
					va := ccr.Get("service.tloc.ip.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocIp = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceVpn = types.Int64Null()

				if t := ccr.Get("service.vpn.optionType"); t.Exists() {
					va := ccr.Get("service.vpn.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceVpn = types.Int64Value(va.Int())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceType = types.StringNull()

				if t := ccr.Get("service.type.optionType"); t.Exists() {
					va := ccr.Get("service.type.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceType = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceLocal = types.BoolNull()

				if t := ccr.Get("service.local.optionType"); t.Exists() {
					va := ccr.Get("service.local.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceLocal = types.BoolValue(va.Bool())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceRestrict = types.BoolNull()

				if t := ccr.Get("service.restrict.optionType"); t.Exists() {
					va := ccr.Get("service.restrict.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceRestrict = types.BoolValue(va.Bool())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocListId = types.StringNull()

				if t := ccr.Get("service.tlocList.refId.optionType"); t.Exists() {
					va := ccr.Get("service.tlocList.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceTlocListId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainType = types.StringNull()

				if t := ccr.Get("serviceChain.type.optionType"); t.Exists() {
					va := ccr.Get("serviceChain.type.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainType = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainVpn = types.Int64Null()

				if t := ccr.Get("serviceChain.vpn.optionType"); t.Exists() {
					va := ccr.Get("serviceChain.vpn.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainVpn = types.Int64Value(va.Int())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainLocal = types.BoolNull()

				if t := ccr.Get("serviceChain.local.optionType"); t.Exists() {
					va := ccr.Get("serviceChain.local.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainLocal = types.BoolValue(va.Bool())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainFallbackToRouting = types.BoolNull()

				if t := ccr.Get("serviceChain.restrict.optionType"); t.Exists() {
					va := ccr.Get("serviceChain.restrict.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainFallbackToRouting = types.BoolValue(va.Bool())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocColor = types.SetNull(types.StringType)

				if t := ccr.Get("serviceChain.tloc.color.optionType"); t.Exists() {
					va := ccr.Get("serviceChain.tloc.color.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocColor = helpers.GetStringSet(va.Array())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocEncapsulation = types.StringNull()

				if t := ccr.Get("serviceChain.tloc.encap.optionType"); t.Exists() {
					va := ccr.Get("serviceChain.tloc.encap.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocEncapsulation = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocIp = types.StringNull()

				if t := ccr.Get("serviceChain.tloc.ip.optionType"); t.Exists() {
					va := ccr.Get("serviceChain.tloc.ip.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocIp = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocListId = types.StringNull()

				if t := ccr.Get("serviceChain.tlocList.refId.optionType"); t.Exists() {
					va := ccr.Get("serviceChain.tlocList.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].ServiceChainTlocListId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].NextHopIpv4 = types.StringNull()

				if t := ccr.Get("nextHop.optionType"); t.Exists() {
					va := ccr.Get("nextHop.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].NextHopIpv4 = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].NextHopIpv6 = types.StringNull()

				if t := ccr.Get("nextHopIpv6.optionType"); t.Exists() {
					va := ccr.Get("nextHopIpv6.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].NextHopIpv6 = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].NextHopLoose = types.BoolNull()

				if t := ccr.Get("nextHopLoose.optionType"); t.Exists() {
					va := ccr.Get("nextHopLoose.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].NextHopLoose = types.BoolValue(va.Bool())
					}
				}
				data.Sequences[i].Actions[ci].SetParameters[cci].Vpn = types.Int64Null()

				if t := ccr.Get("vpn.optionType"); t.Exists() {
					va := ccr.Get("vpn.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SetParameters[cci].Vpn = types.Int64Value(va.Int())
					}
				}
			}
			data.Sequences[i].Actions[ci].RedirectDnsField = types.StringNull()

			if t := cr.Get("redirectDns.field.optionType"); t.Exists() {
				va := cr.Get("redirectDns.field.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].RedirectDnsField = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].RedirectDnsValue = types.StringNull()

			if t := cr.Get("redirectDns.value.optionType"); t.Exists() {
				va := cr.Get("redirectDns.value.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].RedirectDnsValue = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].LossCorrectType = types.StringNull()

			if t := cr.Get("lossCorrection.lossCorrectionType.optionType"); t.Exists() {
				va := cr.Get("lossCorrection.lossCorrectionType.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].LossCorrectType = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].LossCorrectFecThreshold = types.Int64Null()

			if t := cr.Get("lossCorrection.lossCorrectFec.optionType"); t.Exists() {
				va := cr.Get("lossCorrection.lossCorrectFec.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].LossCorrectFecThreshold = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].Actions[ci].Count = types.StringNull()

			if t := cr.Get("count.optionType"); t.Exists() {
				va := cr.Get("count.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].Count = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].Log = types.BoolNull()

			if t := cr.Get("log.optionType"); t.Exists() {
				va := cr.Get("log.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].Log = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].CloudSaas = types.BoolNull()

			if t := cr.Get("cloudSaas.optionType"); t.Exists() {
				va := cr.Get("cloudSaas.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].CloudSaas = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].CloudProbe = types.BoolNull()

			if t := cr.Get("cloudProbe.optionType"); t.Exists() {
				va := cr.Get("cloudProbe.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].CloudProbe = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].NatPool = types.Int64Null()

			if t := cr.Get("natPool.optionType"); t.Exists() {
				va := cr.Get("natPool.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].NatPool = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].Actions[ci].NatVpn = types.BoolNull()

			if t := cr.Get("nat.useVpn.optionType"); t.Exists() {
				va := cr.Get("nat.useVpn.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].NatVpn = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].NatFallback = types.BoolNull()

			if t := cr.Get("nat.fallback.optionType"); t.Exists() {
				va := cr.Get("nat.fallback.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].NatFallback = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].NatBypass = types.BoolNull()

			if t := cr.Get("nat.bypass.optionType"); t.Exists() {
				va := cr.Get("nat.bypass.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].NatBypass = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].NatDiaPool = types.SetNull(types.Int64Type)

			if t := cr.Get("nat.diaPool.optionType"); t.Exists() {
				va := cr.Get("nat.diaPool.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].NatDiaPool = helpers.GetInt64Set(va.Array())
				}
			}
			data.Sequences[i].Actions[ci].NatDiaInterface = types.SetNull(types.StringType)

			if t := cr.Get("nat.diaInterface.optionType"); t.Exists() {
				va := cr.Get("nat.diaInterface.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].NatDiaInterface = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].Actions[ci].SecureInternetGateway = types.BoolNull()

			if t := cr.Get("sig.optionType"); t.Exists() {
				va := cr.Get("sig.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].SecureInternetGateway = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].FallbackToRouting = types.BoolNull()

			if t := cr.Get("fallbackToRouting.optionType"); t.Exists() {
				va := cr.Get("fallbackToRouting.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].FallbackToRouting = types.BoolValue(va.Bool())
				}
			}
		}
	}
}

// End of section. //template:end updateFromBody
