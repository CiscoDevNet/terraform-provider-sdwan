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
type TopologyCustomControl struct {
	Id                    types.String                                 `tfsdk:"id"`
	Version               types.Int64                                  `tfsdk:"version"`
	Name                  types.String                                 `tfsdk:"name"`
	Description           types.String                                 `tfsdk:"description"`
	FeatureProfileId      types.String                                 `tfsdk:"feature_profile_id"`
	DefaultAction         types.String                                 `tfsdk:"default_action"`
	TargetVpn             types.Set                                    `tfsdk:"target_vpn"`
	TargetRole            types.String                                 `tfsdk:"target_role"`
	TargetLevel           types.String                                 `tfsdk:"target_level"`
	TargetInboundSites    types.Set                                    `tfsdk:"target_inbound_sites"`
	TargetOutboundSites   types.Set                                    `tfsdk:"target_outbound_sites"`
	TargetInboundRegions  []TopologyCustomControlTargetInboundRegions  `tfsdk:"target_inbound_regions"`
	TargetOutboundRegions []TopologyCustomControlTargetOutboundRegions `tfsdk:"target_outbound_regions"`
	Sequences             []TopologyCustomControlSequences             `tfsdk:"sequences"`
}

type TopologyCustomControlTargetInboundRegions struct {
	Region     types.String `tfsdk:"region"`
	SubRegions types.Set    `tfsdk:"sub_regions"`
}

type TopologyCustomControlTargetOutboundRegions struct {
	Region     types.String `tfsdk:"region"`
	SubRegions types.Set    `tfsdk:"sub_regions"`
}

type TopologyCustomControlSequences struct {
	Id            types.Int64                                   `tfsdk:"id"`
	Name          types.String                                  `tfsdk:"name"`
	BaseAction    types.String                                  `tfsdk:"base_action"`
	Type          types.String                                  `tfsdk:"type"`
	IpType        types.String                                  `tfsdk:"ip_type"`
	MatchEntries  []TopologyCustomControlSequencesMatchEntries  `tfsdk:"match_entries"`
	ActionEntries []TopologyCustomControlSequencesActionEntries `tfsdk:"action_entries"`
}

type TopologyCustomControlSequencesMatchEntries struct {
	ColorListId             types.String                                             `tfsdk:"color_list_id"`
	CommunityListId         types.String                                             `tfsdk:"community_list_id"`
	ExpandedCommunityListId types.String                                             `tfsdk:"expanded_community_list_id"`
	OmpTag                  types.Int64                                              `tfsdk:"omp_tag"`
	Origin                  types.String                                             `tfsdk:"origin"`
	Originator              types.String                                             `tfsdk:"originator"`
	Preference              types.Int64                                              `tfsdk:"preference"`
	Site                    types.Set                                                `tfsdk:"site"`
	MatchRegions            []TopologyCustomControlSequencesMatchEntriesMatchRegions `tfsdk:"match_regions"`
	Role                    types.String                                             `tfsdk:"role"`
	PathType                types.String                                             `tfsdk:"path_type"`
	TlocListId              types.String                                             `tfsdk:"tloc_list_id"`
	Vpn                     types.Set                                                `tfsdk:"vpn"`
	PrefixListId            types.String                                             `tfsdk:"prefix_list_id"`
	Ipv6PrefixListId        types.String                                             `tfsdk:"ipv6_prefix_list_id"`
	Carrier                 types.String                                             `tfsdk:"carrier"`
	DomainId                types.Int64                                              `tfsdk:"domain_id"`
	GroupId                 types.Int64                                              `tfsdk:"group_id"`
	TlocIp                  types.String                                             `tfsdk:"tloc_ip"`
	TlocColor               types.String                                             `tfsdk:"tloc_color"`
	TlocEncapsulation       types.String                                             `tfsdk:"tloc_encapsulation"`
}
type TopologyCustomControlSequencesActionEntries struct {
	ExportToVpn   types.Set                                                  `tfsdk:"export_to_vpn"`
	SetParameters []TopologyCustomControlSequencesActionEntriesSetParameters `tfsdk:"set_parameters"`
}

type TopologyCustomControlSequencesMatchEntriesMatchRegions struct {
	Region     types.String `tfsdk:"region"`
	SubRegions types.Set    `tfsdk:"sub_regions"`
}

type TopologyCustomControlSequencesActionEntriesSetParameters struct {
	Preference                    types.Int64  `tfsdk:"preference"`
	OmpTag                        types.Int64  `tfsdk:"omp_tag"`
	Community                     types.String `tfsdk:"community"`
	CommunityAdditive             types.Bool   `tfsdk:"community_additive"`
	Affinity                      types.Int64  `tfsdk:"affinity"`
	ServiceType                   types.String `tfsdk:"service_type"`
	ServiceVpn                    types.Int64  `tfsdk:"service_vpn"`
	ServiceTlocIp                 types.String `tfsdk:"service_tloc_ip"`
	ServiceTlocColor              types.String `tfsdk:"service_tloc_color"`
	ServiceTlocEncapsulation      types.String `tfsdk:"service_tloc_encapsulation"`
	ServiceTlocListId             types.String `tfsdk:"service_tloc_list_id"`
	ServiceChainType              types.String `tfsdk:"service_chain_type"`
	ServiceChainVpn               types.Int64  `tfsdk:"service_chain_vpn"`
	ServiceChainTlocIp            types.String `tfsdk:"service_chain_tloc_ip"`
	ServiceChainTlocColor         types.String `tfsdk:"service_chain_tloc_color"`
	ServiceChainTlocEncapsulation types.String `tfsdk:"service_chain_tloc_encapsulation"`
	ServiceChainTlocListId        types.String `tfsdk:"service_chain_tloc_list_id"`
	TlocIp                        types.String `tfsdk:"tloc_ip"`
	TlocColor                     types.String `tfsdk:"tloc_color"`
	TlocEncapsulation             types.String `tfsdk:"tloc_encapsulation"`
	TlocListId                    types.String `tfsdk:"tloc_list_id"`
	TlocAction                    types.String `tfsdk:"tloc_action"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TopologyCustomControl) getModel() string {
	return "topology_custom_control"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TopologyCustomControl) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/topology/%v/custom-control", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TopologyCustomControl) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.DefaultAction.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"defaultAction.optionType", "global")
			body, _ = sjson.Set(body, path+"defaultAction.value", data.DefaultAction.ValueString())
		}
	}
	if !data.TargetVpn.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.vpn.optionType", "global")
			var values []string
			data.TargetVpn.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"target.vpn.value", values)
		}
	}
	if !data.TargetRole.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.role.optionType", "global")
			body, _ = sjson.Set(body, path+"target.role.value", data.TargetRole.ValueString())
		}
	}
	if !data.TargetLevel.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.level.optionType", "global")
			body, _ = sjson.Set(body, path+"target.level.value", data.TargetLevel.ValueString())
		}
	}
	if !data.TargetInboundSites.IsNull() {
		if true && data.TargetLevel.ValueString() == "SITE" {
			body, _ = sjson.Set(body, path+"target.inboundSites.optionType", "global")
			var values []string
			data.TargetInboundSites.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"target.inboundSites.value", values)
		}
	}
	if data.TargetOutboundSites.IsNull() {
		if true && data.TargetLevel.ValueString() == "SITE" {
			body, _ = sjson.Set(body, path+"target.outboundSites.optionType", "global")
			body, _ = sjson.Set(body, path+"target.outboundSites.value", []interface{}{})
		}
	} else {
		if true && data.TargetLevel.ValueString() == "SITE" {
			body, _ = sjson.Set(body, path+"target.outboundSites.optionType", "global")
			var values []string
			data.TargetOutboundSites.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"target.outboundSites.value", values)
		}
	}
	if true && (data.TargetLevel.ValueString() == "REGION" || data.TargetLevel.ValueString() == "SUB_REGION") {

		for _, item := range data.TargetInboundRegions {
			itemBody := ""
			if !item.Region.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "region.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "region.value", item.Region.ValueString())
				}
			}
			if !item.SubRegions.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "subRegion.optionType", "global")
					var values []string
					item.SubRegions.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "subRegion.value", values)
				}
			}
			body, _ = sjson.SetRaw(body, path+"target.inboundRegions.-1", itemBody)
		}
	}
	if true && (data.TargetLevel.ValueString() == "REGION" || data.TargetLevel.ValueString() == "SUB_REGION") {

		for _, item := range data.TargetOutboundRegions {
			itemBody := ""
			if !item.Region.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "region.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "region.value", item.Region.ValueString())
				}
			}
			if !item.SubRegions.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "subRegion.optionType", "global")
					var values []string
					item.SubRegions.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "subRegion.value", values)
				}
			}
			body, _ = sjson.SetRaw(body, path+"target.outboundRegions.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.Id.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceId.value", item.Id.ValueInt64())
				}
			}
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceName.value", item.Name.ValueString())
				}
			}
			if !item.BaseAction.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "baseAction.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "baseAction.value", item.BaseAction.ValueString())
				}
			}
			if !item.Type.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceType.value", item.Type.ValueString())
				}
			}
			if !item.IpType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceIpType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceIpType.value", item.IpType.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.ColorListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "colorList.refId.colorList.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "colorList.refId.colorList.value", childItem.ColorListId.ValueString())
						}
					}
					if !childItem.CommunityListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "community.refId.community.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "community.refId.community.value", childItem.CommunityListId.ValueString())
						}
					}
					if !childItem.ExpandedCommunityListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "expandedCommunity.refId.expandedCommunity.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "expandedCommunity.refId.expandedCommunity.value", childItem.ExpandedCommunityListId.ValueString())
						}
					}
					if !childItem.OmpTag.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ompTag.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ompTag.value", childItem.OmpTag.ValueInt64())
						}
					}
					if !childItem.Origin.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "origin.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "origin.value", childItem.Origin.ValueString())
						}
					}
					if !childItem.Originator.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "originator.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "originator.value", childItem.Originator.ValueString())
						}
					}
					if !childItem.Preference.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "preference.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "preference.value", childItem.Preference.ValueInt64())
						}
					}
					if !childItem.Site.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "site.optionType", "global")
							var values []string
							childItem.Site.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "site.value", values)
						}
					}
					if true {

						for _, childChildItem := range childItem.MatchRegions {
							itemChildChildBody := ""
							if !childChildItem.Region.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "region.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "region.value", childChildItem.Region.ValueString())
								}
							}
							if !childChildItem.SubRegions.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "subRegion.optionType", "global")
									var values []string
									childChildItem.SubRegions.ElementsAs(ctx, &values, false)
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "subRegion.value", values)
								}
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "regions.-1", itemChildChildBody)
						}
					}
					if !childItem.Role.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "role.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "role.value", childItem.Role.ValueString())
						}
					}
					if !childItem.PathType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "pathType.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "pathType.value", childItem.PathType.ValueString())
						}
					}
					if !childItem.TlocListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "tlocList.refId.tlocList.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "tlocList.refId.tlocList.value", childItem.TlocListId.ValueString())
						}
					}
					if !childItem.Vpn.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "vpn.optionType", "global")
							var values []string
							childItem.Vpn.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "vpn.value", values)
						}
					}
					if !childItem.PrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefixList.refId.prefixList.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefixList.refId.prefixList.value", childItem.PrefixListId.ValueString())
						}
					}
					if !childItem.Ipv6PrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6prefixList.refId.ipv6prefixList.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6prefixList.refId.ipv6prefixList.value", childItem.Ipv6PrefixListId.ValueString())
						}
					}
					if !childItem.Carrier.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "carrier.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "carrier.value", childItem.Carrier.ValueString())
						}
					}
					if !childItem.DomainId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "domainId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "domainId.value", childItem.DomainId.ValueInt64())
						}
					}
					if !childItem.GroupId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "groupId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "groupId.value", childItem.GroupId.ValueInt64())
						}
					}
					if !childItem.TlocIp.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "tloc.ip.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "tloc.ip.value", childItem.TlocIp.ValueString())
						}
					}
					if !childItem.TlocColor.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "tloc.color.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "tloc.color.value", childItem.TlocColor.ValueString())
						}
					}
					if !childItem.TlocEncapsulation.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "tloc.encap.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "tloc.encap.value", childItem.TlocEncapsulation.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "actions", []interface{}{})
				for _, childItem := range item.ActionEntries {
					itemChildBody := ""
					if !childItem.ExportToVpn.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "exportTo.optionType", "global")
							var values []string
							childItem.ExportToVpn.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "exportTo.value", values)
						}
					}
					if true {

						for _, childChildItem := range childItem.SetParameters {
							itemChildChildBody := ""
							if !childChildItem.Preference.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preference.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preference.value", childChildItem.Preference.ValueInt64())
								}
							}
							if !childChildItem.OmpTag.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "ompTag.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "ompTag.value", childChildItem.OmpTag.ValueInt64())
								}
							}
							if !childChildItem.Community.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "community.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "community.value", childChildItem.Community.ValueString())
								}
							}
							if !childChildItem.CommunityAdditive.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "communityAdditive.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "communityAdditive.value", childChildItem.CommunityAdditive.ValueBool())
								}
							}
							if !childChildItem.Affinity.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "affinity.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "affinity.value", childChildItem.Affinity.ValueInt64())
								}
							}
							if !childChildItem.ServiceType.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.type.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.type.value", childChildItem.ServiceType.ValueString())
								}
							}
							if !childChildItem.ServiceVpn.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.vpn.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.vpn.value", childChildItem.ServiceVpn.ValueInt64())
								}
							}
							if !childChildItem.ServiceTlocIp.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.ip.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.ip.value", childChildItem.ServiceTlocIp.ValueString())
								}
							}
							if !childChildItem.ServiceTlocColor.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.color.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.color.value", childChildItem.ServiceTlocColor.ValueString())
								}
							}
							if !childChildItem.ServiceTlocEncapsulation.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.encap.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.encap.value", childChildItem.ServiceTlocEncapsulation.ValueString())
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
							if !childChildItem.ServiceChainTlocIp.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.ip.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.ip.value", childChildItem.ServiceChainTlocIp.ValueString())
								}
							}
							if !childChildItem.ServiceChainTlocColor.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.color.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.color.value", childChildItem.ServiceChainTlocColor.ValueString())
								}
							}
							if !childChildItem.ServiceChainTlocEncapsulation.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.encap.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tloc.encap.value", childChildItem.ServiceChainTlocEncapsulation.ValueString())
								}
							}
							if !childChildItem.ServiceChainTlocListId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tlocList.refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "serviceChain.tlocList.refId.value", childChildItem.ServiceChainTlocListId.ValueString())
								}
							}
							if !childChildItem.TlocIp.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.ip.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.ip.value", childChildItem.TlocIp.ValueString())
								}
							}
							if !childChildItem.TlocColor.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.color.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.color.value", childChildItem.TlocColor.ValueString())
								}
							}
							if !childChildItem.TlocEncapsulation.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.encap.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.encap.value", childChildItem.TlocEncapsulation.ValueString())
								}
							}
							if !childChildItem.TlocListId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocList.refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocList.refId.value", childChildItem.TlocListId.ValueString())
								}
							}
							if !childChildItem.TlocAction.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocAction.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocAction.value", childChildItem.TlocAction.ValueString())
								}
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "set.-1", itemChildChildBody)
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
func (data *TopologyCustomControl) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "defaultAction.optionType"); t.Exists() {
		va := res.Get(path + "defaultAction.value")
		if t.String() == "global" {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	data.TargetVpn = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			data.TargetVpn = helpers.GetStringSet(va.Array())
		}
	}
	data.TargetRole = types.StringNull()

	if t := res.Get(path + "target.role.optionType"); t.Exists() {
		va := res.Get(path + "target.role.value")
		if t.String() == "global" {
			data.TargetRole = types.StringValue(va.String())
		}
	}
	data.TargetLevel = types.StringNull()

	if t := res.Get(path + "target.level.optionType"); t.Exists() {
		va := res.Get(path + "target.level.value")
		if t.String() == "global" {
			data.TargetLevel = types.StringValue(va.String())
		}
	}
	data.TargetInboundSites = types.SetNull(types.StringType)

	if t := res.Get(path + "target.inboundSites.optionType"); t.Exists() {
		va := res.Get(path + "target.inboundSites.value")
		if t.String() == "global" {
			data.TargetInboundSites = helpers.GetStringSet(va.Array())
		}
	}
	data.TargetOutboundSites = types.SetNull(types.StringType)

	if t := res.Get(path + "target.outboundSites.optionType"); t.Exists() {
		va := res.Get(path + "target.outboundSites.value")
		if t.String() == "global" {
			data.TargetOutboundSites = helpers.GetStringSet(va.Array())
		}
	}
	if value := res.Get(path + "target.inboundRegions"); value.Exists() && len(value.Array()) > 0 {
		data.TargetInboundRegions = make([]TopologyCustomControlTargetInboundRegions, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TopologyCustomControlTargetInboundRegions{}
			item.Region = types.StringNull()

			if t := v.Get("region.optionType"); t.Exists() {
				va := v.Get("region.value")
				if t.String() == "global" {
					item.Region = types.StringValue(va.String())
				}
			}
			item.SubRegions = types.SetNull(types.StringType)

			if t := v.Get("subRegion.optionType"); t.Exists() {
				va := v.Get("subRegion.value")
				if t.String() == "global" {
					item.SubRegions = helpers.GetStringSet(va.Array())
				}
			}
			data.TargetInboundRegions = append(data.TargetInboundRegions, item)
			return true
		})
	}
	if value := res.Get(path + "target.outboundRegions"); value.Exists() && len(value.Array()) > 0 {
		data.TargetOutboundRegions = make([]TopologyCustomControlTargetOutboundRegions, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TopologyCustomControlTargetOutboundRegions{}
			item.Region = types.StringNull()

			if t := v.Get("region.optionType"); t.Exists() {
				va := v.Get("region.value")
				if t.String() == "global" {
					item.Region = types.StringValue(va.String())
				}
			}
			item.SubRegions = types.SetNull(types.StringType)

			if t := v.Get("subRegion.optionType"); t.Exists() {
				va := v.Get("subRegion.value")
				if t.String() == "global" {
					item.SubRegions = helpers.GetStringSet(va.Array())
				}
			}
			data.TargetOutboundRegions = append(data.TargetOutboundRegions, item)
			return true
		})
	}
	if value := res.Get(path + "sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]TopologyCustomControlSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TopologyCustomControlSequences{}
			item.Id = types.Int64Null()

			if t := v.Get("sequenceId.optionType"); t.Exists() {
				va := v.Get("sequenceId.value")
				if t.String() == "global" {
					item.Id = types.Int64Value(va.Int())
				}
			}
			item.Name = types.StringNull()

			if t := v.Get("sequenceName.optionType"); t.Exists() {
				va := v.Get("sequenceName.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.BaseAction = types.StringNull()

			if t := v.Get("baseAction.optionType"); t.Exists() {
				va := v.Get("baseAction.value")
				if t.String() == "global" {
					item.BaseAction = types.StringValue(va.String())
				}
			}
			item.Type = types.StringNull()

			if t := v.Get("sequenceType.optionType"); t.Exists() {
				va := v.Get("sequenceType.value")
				if t.String() == "global" {
					item.Type = types.StringValue(va.String())
				}
			}
			item.IpType = types.StringNull()

			if t := v.Get("sequenceIpType.optionType"); t.Exists() {
				va := v.Get("sequenceIpType.value")
				if t.String() == "global" {
					item.IpType = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]TopologyCustomControlSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TopologyCustomControlSequencesMatchEntries{}
					cItem.ColorListId = types.StringNull()

					if t := cv.Get("colorList.refId.colorList.optionType"); t.Exists() {
						va := cv.Get("colorList.refId.colorList.value")
						if t.String() == "global" {
							cItem.ColorListId = types.StringValue(va.String())
						}
					}
					cItem.CommunityListId = types.StringNull()

					if t := cv.Get("community.refId.community.optionType"); t.Exists() {
						va := cv.Get("community.refId.community.value")
						if t.String() == "global" {
							cItem.CommunityListId = types.StringValue(va.String())
						}
					}
					cItem.ExpandedCommunityListId = types.StringNull()

					if t := cv.Get("expandedCommunity.refId.expandedCommunity.optionType"); t.Exists() {
						va := cv.Get("expandedCommunity.refId.expandedCommunity.value")
						if t.String() == "global" {
							cItem.ExpandedCommunityListId = types.StringValue(va.String())
						}
					}
					cItem.OmpTag = types.Int64Null()

					if t := cv.Get("ompTag.optionType"); t.Exists() {
						va := cv.Get("ompTag.value")
						if t.String() == "global" {
							cItem.OmpTag = types.Int64Value(va.Int())
						}
					}
					cItem.Origin = types.StringNull()

					if t := cv.Get("origin.optionType"); t.Exists() {
						va := cv.Get("origin.value")
						if t.String() == "global" {
							cItem.Origin = types.StringValue(va.String())
						}
					}
					cItem.Originator = types.StringNull()

					if t := cv.Get("originator.optionType"); t.Exists() {
						va := cv.Get("originator.value")
						if t.String() == "global" {
							cItem.Originator = types.StringValue(va.String())
						}
					}
					cItem.Preference = types.Int64Null()

					if t := cv.Get("preference.optionType"); t.Exists() {
						va := cv.Get("preference.value")
						if t.String() == "global" {
							cItem.Preference = types.Int64Value(va.Int())
						}
					}
					cItem.Site = types.SetNull(types.StringType)

					if t := cv.Get("site.optionType"); t.Exists() {
						va := cv.Get("site.value")
						if t.String() == "global" {
							cItem.Site = helpers.GetStringSet(va.Array())
						}
					}
					if ccValue := cv.Get("regions"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.MatchRegions = make([]TopologyCustomControlSequencesMatchEntriesMatchRegions, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := TopologyCustomControlSequencesMatchEntriesMatchRegions{}
							ccItem.Region = types.StringNull()

							if t := ccv.Get("region.optionType"); t.Exists() {
								va := ccv.Get("region.value")
								if t.String() == "global" {
									ccItem.Region = types.StringValue(va.String())
								}
							}
							ccItem.SubRegions = types.SetNull(types.StringType)

							if t := ccv.Get("subRegion.optionType"); t.Exists() {
								va := ccv.Get("subRegion.value")
								if t.String() == "global" {
									ccItem.SubRegions = helpers.GetStringSet(va.Array())
								}
							}
							cItem.MatchRegions = append(cItem.MatchRegions, ccItem)
							return true
						})
					}
					cItem.Role = types.StringNull()

					if t := cv.Get("role.optionType"); t.Exists() {
						va := cv.Get("role.value")
						if t.String() == "global" {
							cItem.Role = types.StringValue(va.String())
						}
					}
					cItem.PathType = types.StringNull()

					if t := cv.Get("pathType.optionType"); t.Exists() {
						va := cv.Get("pathType.value")
						if t.String() == "global" {
							cItem.PathType = types.StringValue(va.String())
						}
					}
					cItem.TlocListId = types.StringNull()

					if t := cv.Get("tlocList.refId.tlocList.optionType"); t.Exists() {
						va := cv.Get("tlocList.refId.tlocList.value")
						if t.String() == "global" {
							cItem.TlocListId = types.StringValue(va.String())
						}
					}
					cItem.Vpn = types.SetNull(types.StringType)

					if t := cv.Get("vpn.optionType"); t.Exists() {
						va := cv.Get("vpn.value")
						if t.String() == "global" {
							cItem.Vpn = helpers.GetStringSet(va.Array())
						}
					}
					cItem.PrefixListId = types.StringNull()

					if t := cv.Get("prefixList.refId.prefixList.optionType"); t.Exists() {
						va := cv.Get("prefixList.refId.prefixList.value")
						if t.String() == "global" {
							cItem.PrefixListId = types.StringValue(va.String())
						}
					}
					cItem.Ipv6PrefixListId = types.StringNull()

					if t := cv.Get("ipv6prefixList.refId.ipv6prefixList.optionType"); t.Exists() {
						va := cv.Get("ipv6prefixList.refId.ipv6prefixList.value")
						if t.String() == "global" {
							cItem.Ipv6PrefixListId = types.StringValue(va.String())
						}
					}
					cItem.Carrier = types.StringNull()

					if t := cv.Get("carrier.optionType"); t.Exists() {
						va := cv.Get("carrier.value")
						if t.String() == "global" {
							cItem.Carrier = types.StringValue(va.String())
						}
					}
					cItem.DomainId = types.Int64Null()

					if t := cv.Get("domainId.optionType"); t.Exists() {
						va := cv.Get("domainId.value")
						if t.String() == "global" {
							cItem.DomainId = types.Int64Value(va.Int())
						}
					}
					cItem.GroupId = types.Int64Null()

					if t := cv.Get("groupId.optionType"); t.Exists() {
						va := cv.Get("groupId.value")
						if t.String() == "global" {
							cItem.GroupId = types.Int64Value(va.Int())
						}
					}
					cItem.TlocIp = types.StringNull()

					if t := cv.Get("tloc.ip.optionType"); t.Exists() {
						va := cv.Get("tloc.ip.value")
						if t.String() == "global" {
							cItem.TlocIp = types.StringValue(va.String())
						}
					}
					cItem.TlocColor = types.StringNull()

					if t := cv.Get("tloc.color.optionType"); t.Exists() {
						va := cv.Get("tloc.color.value")
						if t.String() == "global" {
							cItem.TlocColor = types.StringValue(va.String())
						}
					}
					cItem.TlocEncapsulation = types.StringNull()

					if t := cv.Get("tloc.encap.optionType"); t.Exists() {
						va := cv.Get("tloc.encap.value")
						if t.String() == "global" {
							cItem.TlocEncapsulation = types.StringValue(va.String())
						}
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ActionEntries = make([]TopologyCustomControlSequencesActionEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TopologyCustomControlSequencesActionEntries{}
					cItem.ExportToVpn = types.SetNull(types.StringType)

					if t := cv.Get("exportTo.optionType"); t.Exists() {
						va := cv.Get("exportTo.value")
						if t.String() == "global" {
							cItem.ExportToVpn = helpers.GetStringSet(va.Array())
						}
					}
					if ccValue := cv.Get("set"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.SetParameters = make([]TopologyCustomControlSequencesActionEntriesSetParameters, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := TopologyCustomControlSequencesActionEntriesSetParameters{}
							ccItem.Preference = types.Int64Null()

							if t := ccv.Get("preference.optionType"); t.Exists() {
								va := ccv.Get("preference.value")
								if t.String() == "global" {
									ccItem.Preference = types.Int64Value(va.Int())
								}
							}
							ccItem.OmpTag = types.Int64Null()

							if t := ccv.Get("ompTag.optionType"); t.Exists() {
								va := ccv.Get("ompTag.value")
								if t.String() == "global" {
									ccItem.OmpTag = types.Int64Value(va.Int())
								}
							}
							ccItem.Community = types.StringNull()

							if t := ccv.Get("community.optionType"); t.Exists() {
								va := ccv.Get("community.value")
								if t.String() == "global" {
									ccItem.Community = types.StringValue(va.String())
								}
							}
							ccItem.CommunityAdditive = types.BoolNull()

							if t := ccv.Get("communityAdditive.optionType"); t.Exists() {
								va := ccv.Get("communityAdditive.value")
								if t.String() == "global" {
									ccItem.CommunityAdditive = types.BoolValue(va.Bool())
								}
							}
							ccItem.Affinity = types.Int64Null()

							if t := ccv.Get("affinity.optionType"); t.Exists() {
								va := ccv.Get("affinity.value")
								if t.String() == "global" {
									ccItem.Affinity = types.Int64Value(va.Int())
								}
							}
							ccItem.ServiceType = types.StringNull()

							if t := ccv.Get("service.type.optionType"); t.Exists() {
								va := ccv.Get("service.type.value")
								if t.String() == "global" {
									ccItem.ServiceType = types.StringValue(va.String())
								}
							}
							ccItem.ServiceVpn = types.Int64Null()

							if t := ccv.Get("service.vpn.optionType"); t.Exists() {
								va := ccv.Get("service.vpn.value")
								if t.String() == "global" {
									ccItem.ServiceVpn = types.Int64Value(va.Int())
								}
							}
							ccItem.ServiceTlocIp = types.StringNull()

							if t := ccv.Get("service.tloc.ip.optionType"); t.Exists() {
								va := ccv.Get("service.tloc.ip.value")
								if t.String() == "global" {
									ccItem.ServiceTlocIp = types.StringValue(va.String())
								}
							}
							ccItem.ServiceTlocColor = types.StringNull()

							if t := ccv.Get("service.tloc.color.optionType"); t.Exists() {
								va := ccv.Get("service.tloc.color.value")
								if t.String() == "global" {
									ccItem.ServiceTlocColor = types.StringValue(va.String())
								}
							}
							ccItem.ServiceTlocEncapsulation = types.StringNull()

							if t := ccv.Get("service.tloc.encap.optionType"); t.Exists() {
								va := ccv.Get("service.tloc.encap.value")
								if t.String() == "global" {
									ccItem.ServiceTlocEncapsulation = types.StringValue(va.String())
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
							ccItem.ServiceChainTlocIp = types.StringNull()

							if t := ccv.Get("serviceChain.tloc.ip.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.tloc.ip.value")
								if t.String() == "global" {
									ccItem.ServiceChainTlocIp = types.StringValue(va.String())
								}
							}
							ccItem.ServiceChainTlocColor = types.StringNull()

							if t := ccv.Get("serviceChain.tloc.color.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.tloc.color.value")
								if t.String() == "global" {
									ccItem.ServiceChainTlocColor = types.StringValue(va.String())
								}
							}
							ccItem.ServiceChainTlocEncapsulation = types.StringNull()

							if t := ccv.Get("serviceChain.tloc.encap.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.tloc.encap.value")
								if t.String() == "global" {
									ccItem.ServiceChainTlocEncapsulation = types.StringValue(va.String())
								}
							}
							ccItem.ServiceChainTlocListId = types.StringNull()

							if t := ccv.Get("serviceChain.tlocList.refId.optionType"); t.Exists() {
								va := ccv.Get("serviceChain.tlocList.refId.value")
								if t.String() == "global" {
									ccItem.ServiceChainTlocListId = types.StringValue(va.String())
								}
							}
							ccItem.TlocIp = types.StringNull()

							if t := ccv.Get("tloc.ip.optionType"); t.Exists() {
								va := ccv.Get("tloc.ip.value")
								if t.String() == "global" {
									ccItem.TlocIp = types.StringValue(va.String())
								}
							}
							ccItem.TlocColor = types.StringNull()

							if t := ccv.Get("tloc.color.optionType"); t.Exists() {
								va := ccv.Get("tloc.color.value")
								if t.String() == "global" {
									ccItem.TlocColor = types.StringValue(va.String())
								}
							}
							ccItem.TlocEncapsulation = types.StringNull()

							if t := ccv.Get("tloc.encap.optionType"); t.Exists() {
								va := ccv.Get("tloc.encap.value")
								if t.String() == "global" {
									ccItem.TlocEncapsulation = types.StringValue(va.String())
								}
							}
							ccItem.TlocListId = types.StringNull()

							if t := ccv.Get("tlocList.refId.optionType"); t.Exists() {
								va := ccv.Get("tlocList.refId.value")
								if t.String() == "global" {
									ccItem.TlocListId = types.StringValue(va.String())
								}
							}
							ccItem.TlocAction = types.StringNull()

							if t := ccv.Get("tlocAction.optionType"); t.Exists() {
								va := ccv.Get("tlocAction.value")
								if t.String() == "global" {
									ccItem.TlocAction = types.StringValue(va.String())
								}
							}
							cItem.SetParameters = append(cItem.SetParameters, ccItem)
							return true
						})
					}
					item.ActionEntries = append(item.ActionEntries, cItem)
					return true
				})
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

func (data *TopologyCustomControl) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "defaultAction.optionType"); t.Exists() {
		va := res.Get(path + "defaultAction.value")
		if t.String() == "global" {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	data.TargetVpn = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			if len(va.Array()) > 0 {
				data.TargetVpn = helpers.GetStringSet(va.Array())
			}
		}
	}
	data.TargetRole = types.StringNull()

	if t := res.Get(path + "target.role.optionType"); t.Exists() {
		va := res.Get(path + "target.role.value")
		if t.String() == "global" {
			data.TargetRole = types.StringValue(va.String())
		}
	}
	data.TargetLevel = types.StringNull()

	if t := res.Get(path + "target.level.optionType"); t.Exists() {
		va := res.Get(path + "target.level.value")
		if t.String() == "global" {
			data.TargetLevel = types.StringValue(va.String())
		}
	}
	data.TargetInboundSites = types.SetNull(types.StringType)

	if t := res.Get(path + "target.inboundSites.optionType"); t.Exists() {
		va := res.Get(path + "target.inboundSites.value")
		if t.String() == "global" {
			if len(va.Array()) > 0 {
				data.TargetInboundSites = helpers.GetStringSet(va.Array())
			}
		}
	}
	data.TargetOutboundSites = types.SetNull(types.StringType)

	if t := res.Get(path + "target.outboundSites.optionType"); t.Exists() {
		va := res.Get(path + "target.outboundSites.value")
		if t.String() == "global" {
			if len(va.Array()) > 0 {
				data.TargetOutboundSites = helpers.GetStringSet(va.Array())
			}
		}
	}
	for i := range data.TargetInboundRegions {
		keys := [...]string{"region", "subRegion"}
		keyValues := [...]string{data.TargetInboundRegions[i].Region.ValueString(), helpers.GetStringFromSet(data.TargetInboundRegions[i].SubRegions).ValueString()}
		keyValuesVariables := [...]string{"", ""}

		var r gjson.Result
		res.Get(path + "target.inboundRegions").ForEach(
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
		data.TargetInboundRegions[i].Region = types.StringNull()

		if t := r.Get("region.optionType"); t.Exists() {
			va := r.Get("region.value")
			if t.String() == "global" {
				data.TargetInboundRegions[i].Region = types.StringValue(va.String())
			}
		}
		data.TargetInboundRegions[i].SubRegions = types.SetNull(types.StringType)

		if t := r.Get("subRegion.optionType"); t.Exists() {
			va := r.Get("subRegion.value")
			if t.String() == "global" {
				data.TargetInboundRegions[i].SubRegions = helpers.GetStringSet(va.Array())
			}
		}
	}
	for i := range data.TargetOutboundRegions {
		keys := [...]string{"region", "subRegion"}
		keyValues := [...]string{data.TargetOutboundRegions[i].Region.ValueString(), helpers.GetStringFromSet(data.TargetOutboundRegions[i].SubRegions).ValueString()}
		keyValuesVariables := [...]string{"", ""}

		var r gjson.Result
		res.Get(path + "target.outboundRegions").ForEach(
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
		data.TargetOutboundRegions[i].Region = types.StringNull()

		if t := r.Get("region.optionType"); t.Exists() {
			va := r.Get("region.value")
			if t.String() == "global" {
				data.TargetOutboundRegions[i].Region = types.StringValue(va.String())
			}
		}
		data.TargetOutboundRegions[i].SubRegions = types.SetNull(types.StringType)

		if t := r.Get("subRegion.optionType"); t.Exists() {
			va := r.Get("subRegion.value")
			if t.String() == "global" {
				data.TargetOutboundRegions[i].SubRegions = helpers.GetStringSet(va.Array())
			}
		}
	}
	for i := range data.Sequences {
		keys := [...]string{"sequenceId"}
		keyValues := [...]string{strconv.FormatInt(data.Sequences[i].Id.ValueInt64(), 10)}
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
		if !r.Exists() {
			arr := res.Get(path + "sequences").Array()
			if i < len(arr) {
				r = arr[i]
			}
		}
		data.Sequences[i].Id = types.Int64Null()

		if t := r.Get("sequenceId.optionType"); t.Exists() {
			va := r.Get("sequenceId.value")
			if t.String() == "global" {
				data.Sequences[i].Id = types.Int64Value(va.Int())
			}
		}
		data.Sequences[i].Name = types.StringNull()

		if t := r.Get("sequenceName.optionType"); t.Exists() {
			va := r.Get("sequenceName.value")
			if t.String() == "global" {
				data.Sequences[i].Name = types.StringValue(va.String())
			}
		}
		data.Sequences[i].BaseAction = types.StringNull()

		if t := r.Get("baseAction.optionType"); t.Exists() {
			va := r.Get("baseAction.value")
			if t.String() == "global" {
				data.Sequences[i].BaseAction = types.StringValue(va.String())
			}
		}
		data.Sequences[i].Type = types.StringNull()

		if t := r.Get("sequenceType.optionType"); t.Exists() {
			va := r.Get("sequenceType.value")
			if t.String() == "global" {
				data.Sequences[i].Type = types.StringValue(va.String())
			}
		}
		data.Sequences[i].IpType = types.StringNull()

		if t := r.Get("sequenceIpType.optionType"); t.Exists() {
			va := r.Get("sequenceIpType.value")
			if t.String() == "global" {
				data.Sequences[i].IpType = types.StringValue(va.String())
			}
		}
		matchEntriesArray := r.Get("match.entries").Array()
		matchEntriesByKey := make(map[string]gjson.Result)
		for _, elem := range matchEntriesArray {
			elem.ForEach(func(key, value gjson.Result) bool {
				matchEntriesByKey[key.String()] = elem
				return true
			})
		}
		for ci := range data.Sequences[i].MatchEntries {
			found := false
			var cr gjson.Result
			if !data.Sequences[i].MatchEntries[ci].ColorListId.IsNull() {
				if elem, ok := matchEntriesByKey["colorList"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].CommunityListId.IsNull() {
				if elem, ok := matchEntriesByKey["community"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].ExpandedCommunityListId.IsNull() {
				if elem, ok := matchEntriesByKey["expandedCommunity"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].OmpTag.IsNull() {
				if elem, ok := matchEntriesByKey["ompTag"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].Origin.IsNull() {
				if elem, ok := matchEntriesByKey["origin"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].Originator.IsNull() {
				if elem, ok := matchEntriesByKey["originator"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].Preference.IsNull() {
				if elem, ok := matchEntriesByKey["preference"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].Site.IsNull() {
				if elem, ok := matchEntriesByKey["site"]; ok {
					cr = elem
					found = true
				}
			} else if len(data.Sequences[i].MatchEntries[ci].MatchRegions) > 0 {
				if elem, ok := matchEntriesByKey["regions"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].Role.IsNull() {
				if elem, ok := matchEntriesByKey["role"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].PathType.IsNull() {
				if elem, ok := matchEntriesByKey["pathType"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].TlocListId.IsNull() {
				if elem, ok := matchEntriesByKey["tlocList"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].Vpn.IsNull() {
				if elem, ok := matchEntriesByKey["vpn"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].PrefixListId.IsNull() {
				if elem, ok := matchEntriesByKey["prefixList"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].Ipv6PrefixListId.IsNull() {
				if elem, ok := matchEntriesByKey["ipv6prefixList"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].Carrier.IsNull() {
				if elem, ok := matchEntriesByKey["carrier"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].DomainId.IsNull() {
				if elem, ok := matchEntriesByKey["domainId"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].GroupId.IsNull() {
				if elem, ok := matchEntriesByKey["groupId"]; ok {
					cr = elem
					found = true
				}
			} else if !data.Sequences[i].MatchEntries[ci].TlocIp.IsNull() || !data.Sequences[i].MatchEntries[ci].TlocColor.IsNull() || !data.Sequences[i].MatchEntries[ci].TlocEncapsulation.IsNull() {
				if elem, ok := matchEntriesByKey["tloc"]; ok {
					cr = elem
					found = true
				}
			}
			if !found {
				if ci < len(matchEntriesArray) {
					cr = matchEntriesArray[ci]
				}
			}
			data.Sequences[i].MatchEntries[ci].ColorListId = types.StringNull()
			data.Sequences[i].MatchEntries[ci].CommunityListId = types.StringNull()
			data.Sequences[i].MatchEntries[ci].ExpandedCommunityListId = types.StringNull()
			data.Sequences[i].MatchEntries[ci].OmpTag = types.Int64Null()
			data.Sequences[i].MatchEntries[ci].Origin = types.StringNull()
			data.Sequences[i].MatchEntries[ci].Originator = types.StringNull()
			data.Sequences[i].MatchEntries[ci].Preference = types.Int64Null()
			data.Sequences[i].MatchEntries[ci].Site = types.SetNull(types.StringType)
			data.Sequences[i].MatchEntries[ci].Role = types.StringNull()
			data.Sequences[i].MatchEntries[ci].PathType = types.StringNull()
			data.Sequences[i].MatchEntries[ci].TlocListId = types.StringNull()
			data.Sequences[i].MatchEntries[ci].Vpn = types.SetNull(types.StringType)
			data.Sequences[i].MatchEntries[ci].PrefixListId = types.StringNull()
			data.Sequences[i].MatchEntries[ci].Ipv6PrefixListId = types.StringNull()
			data.Sequences[i].MatchEntries[ci].Carrier = types.StringNull()
			data.Sequences[i].MatchEntries[ci].DomainId = types.Int64Null()
			data.Sequences[i].MatchEntries[ci].GroupId = types.Int64Null()
			data.Sequences[i].MatchEntries[ci].TlocIp = types.StringNull()
			data.Sequences[i].MatchEntries[ci].TlocColor = types.StringNull()
			data.Sequences[i].MatchEntries[ci].TlocEncapsulation = types.StringNull()
			if t := cr.Get("colorList.refId.colorList.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].ColorListId = types.StringValue(cr.Get("colorList.refId.colorList.value").String())
			}
			if t := cr.Get("community.refId.community.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].CommunityListId = types.StringValue(cr.Get("community.refId.community.value").String())
			}
			if t := cr.Get("expandedCommunity.refId.expandedCommunity.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].ExpandedCommunityListId = types.StringValue(cr.Get("expandedCommunity.refId.expandedCommunity.value").String())
			}
			if t := cr.Get("ompTag.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].OmpTag = types.Int64Value(cr.Get("ompTag.value").Int())
			}
			if t := cr.Get("origin.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].Origin = types.StringValue(cr.Get("origin.value").String())
			}
			if t := cr.Get("originator.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].Originator = types.StringValue(cr.Get("originator.value").String())
			}
			if t := cr.Get("preference.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].Preference = types.Int64Value(cr.Get("preference.value").Int())
			}
			if t := cr.Get("site.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].Site = helpers.GetStringSet(cr.Get("site.value").Array())
			}
			for cci := range data.Sequences[i].MatchEntries[ci].MatchRegions {
				keys := [...]string{"region", "subRegion"}
				keyValues := [...]string{data.Sequences[i].MatchEntries[ci].MatchRegions[cci].Region.ValueString(), helpers.GetStringFromSet(data.Sequences[i].MatchEntries[ci].MatchRegions[cci].SubRegions).ValueString()}
				keyValuesVariables := [...]string{"", ""}
				var ccr gjson.Result
				cr.Get("regions").ForEach(
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
				data.Sequences[i].MatchEntries[ci].MatchRegions[cci].Region = types.StringNull()
				data.Sequences[i].MatchEntries[ci].MatchRegions[cci].SubRegions = types.SetNull(types.StringType)
				if t := ccr.Get("region.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].MatchRegions[cci].Region = types.StringValue(ccr.Get("region.value").String())
				}
				if t := ccr.Get("subRegion.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].MatchRegions[cci].SubRegions = helpers.GetStringSet(ccr.Get("subRegion.value").Array())
				}
			}
			if t := cr.Get("role.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].Role = types.StringValue(cr.Get("role.value").String())
			}
			if t := cr.Get("pathType.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].PathType = types.StringValue(cr.Get("pathType.value").String())
			}
			if t := cr.Get("tlocList.refId.tlocList.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].TlocListId = types.StringValue(cr.Get("tlocList.refId.tlocList.value").String())
			}
			if t := cr.Get("vpn.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].Vpn = helpers.GetStringSet(cr.Get("vpn.value").Array())
			}
			if t := cr.Get("prefixList.refId.prefixList.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].PrefixListId = types.StringValue(cr.Get("prefixList.refId.prefixList.value").String())
			}
			if t := cr.Get("ipv6prefixList.refId.ipv6prefixList.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].Ipv6PrefixListId = types.StringValue(cr.Get("ipv6prefixList.refId.ipv6prefixList.value").String())
			}
			if t := cr.Get("carrier.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].Carrier = types.StringValue(cr.Get("carrier.value").String())
			}
			if t := cr.Get("domainId.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].DomainId = types.Int64Value(cr.Get("domainId.value").Int())
			}
			if t := cr.Get("groupId.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].GroupId = types.Int64Value(cr.Get("groupId.value").Int())
			}
			if t := cr.Get("tloc.ip.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].TlocIp = types.StringValue(cr.Get("tloc.ip.value").String())
			}
			if t := cr.Get("tloc.color.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].TlocColor = types.StringValue(cr.Get("tloc.color.value").String())
			}
			if t := cr.Get("tloc.encap.optionType"); t.Exists() && t.String() == "global" {
				data.Sequences[i].MatchEntries[ci].TlocEncapsulation = types.StringValue(cr.Get("tloc.encap.value").String())
			}
		}
		// API returns split objects in actions array: [{set:[...]}, {exportTo:{...}}]
		// Build a map from top-level key to the API element containing it
		actionEntriesArray := r.Get("actions").Array()
		fmt.Printf("DEBUG updateFromBody: API actions = %s\n", r.Get("actions").String())
		fmt.Printf("DEBUG updateFromBody: TF ActionEntries count = %d\n", len(data.Sequences[i].ActionEntries))
		actionEntriesByKey := make(map[string]gjson.Result)
		for _, elem := range actionEntriesArray {
			elem.ForEach(func(key, value gjson.Result) bool {
				actionEntriesByKey[key.String()] = elem
				return true
			})
		}
		for ci := range data.Sequences[i].ActionEntries {
			hasExportToVpn := !data.Sequences[i].ActionEntries[ci].ExportToVpn.IsNull()
			hasSetParameters := len(data.Sequences[i].ActionEntries[ci].SetParameters) > 0

			data.Sequences[i].ActionEntries[ci].ExportToVpn = types.SetNull(types.StringType)
			if hasExportToVpn || !hasSetParameters {
				if exportToElem, ok := actionEntriesByKey["exportTo"]; ok {
					if t := exportToElem.Get("exportTo.optionType"); t.Exists() && t.String() == "global" {
						data.Sequences[i].ActionEntries[ci].ExportToVpn = helpers.GetStringSet(exportToElem.Get("exportTo.value").Array())
					}
				}
			}
			var setParamsArray []gjson.Result
			if hasSetParameters {
				if setElem, ok := actionEntriesByKey["set"]; ok {
					setParamsArray = setElem.Get("set").Array()
				}
			}
			setParamsByKey := make(map[string]gjson.Result)
			for _, elem := range setParamsArray {
				elem.ForEach(func(key, value gjson.Result) bool {
					setParamsByKey[key.String()] = elem
					return true
				})
			}
			for cci := range data.Sequences[i].ActionEntries[ci].SetParameters {
				keys := [...]string{"preference", "ompTag", "community", "communityAdditive", "affinity", "service", "serviceChain", "tloc", "tlocList", "tlocAction"}
				found := false
				var ccrr gjson.Result
				for _, k := range keys {
					if !data.Sequences[i].ActionEntries[ci].SetParameters[cci].Preference.IsNull() && k == "preference" {
						if elem, ok := setParamsByKey["preference"]; ok {
							ccrr = elem
							found = true
						}
						break
					}
					if !data.Sequences[i].ActionEntries[ci].SetParameters[cci].OmpTag.IsNull() && k == "ompTag" {
						if elem, ok := setParamsByKey["ompTag"]; ok {
							ccrr = elem
							found = true
						}
						break
					}
					if !data.Sequences[i].ActionEntries[ci].SetParameters[cci].Community.IsNull() && k == "community" {
						if elem, ok := setParamsByKey["community"]; ok {
							ccrr = elem
							found = true
						}
						break
					}
					if !data.Sequences[i].ActionEntries[ci].SetParameters[cci].CommunityAdditive.IsNull() && k == "communityAdditive" {
						if elem, ok := setParamsByKey["communityAdditive"]; ok {
							ccrr = elem
							found = true
						}
						break
					}
					if !data.Sequences[i].ActionEntries[ci].SetParameters[cci].Affinity.IsNull() && k == "affinity" {
						if elem, ok := setParamsByKey["affinity"]; ok {
							ccrr = elem
							found = true
						}
						break
					}
					if (!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceType.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceVpn.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocIp.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocColor.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocEncapsulation.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocListId.IsNull()) && k == "service" {
						if elem, ok := setParamsByKey["service"]; ok {
							ccrr = elem
							found = true
						}
						break
					}
					if (!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainType.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainVpn.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocIp.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocColor.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocEncapsulation.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocListId.IsNull()) && k == "serviceChain" {
						if elem, ok := setParamsByKey["serviceChain"]; ok {
							ccrr = elem
							found = true
						}
						break
					}
					if (!data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocIp.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocColor.IsNull() ||
						!data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocEncapsulation.IsNull()) && k == "tloc" {
						if elem, ok := setParamsByKey["tloc"]; ok {
							ccrr = elem
							found = true
						}
						break
					}
					if !data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocListId.IsNull() && k == "tlocList" {
						if elem, ok := setParamsByKey["tlocList"]; ok {
							ccrr = elem
							found = true
						}
						break
					}
					if !data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocAction.IsNull() && k == "tlocAction" {
						if elem, ok := setParamsByKey["tlocAction"]; ok {
							ccrr = elem
							found = true
						}
						break
					}
				}
				if !found {
					if cci < len(setParamsArray) {
						ccrr = setParamsArray[cci]
					}
				}
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].Preference = types.Int64Null()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].OmpTag = types.Int64Null()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].Community = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].CommunityAdditive = types.BoolNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].Affinity = types.Int64Null()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceType = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceVpn = types.Int64Null()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocIp = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocColor = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocEncapsulation = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocListId = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainType = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainVpn = types.Int64Null()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocIp = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocColor = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocEncapsulation = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocListId = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocIp = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocColor = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocEncapsulation = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocListId = types.StringNull()
				data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocAction = types.StringNull()
				if t := ccrr.Get("preference.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].Preference = types.Int64Value(ccrr.Get("preference.value").Int())
				}
				if t := ccrr.Get("ompTag.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].OmpTag = types.Int64Value(ccrr.Get("ompTag.value").Int())
				}
				if t := ccrr.Get("community.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].Community = types.StringValue(ccrr.Get("community.value").String())
				}
				if t := ccrr.Get("communityAdditive.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].CommunityAdditive = types.BoolValue(ccrr.Get("communityAdditive.value").Bool())
				}
				if t := ccrr.Get("affinity.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].Affinity = types.Int64Value(ccrr.Get("affinity.value").Int())
				}
				if t := ccrr.Get("service.type.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceType = types.StringValue(ccrr.Get("service.type.value").String())
				}
				if t := ccrr.Get("service.vpn.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceVpn = types.Int64Value(ccrr.Get("service.vpn.value").Int())
				}
				if t := ccrr.Get("service.tloc.ip.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocIp = types.StringValue(ccrr.Get("service.tloc.ip.value").String())
				}
				if t := ccrr.Get("service.tloc.color.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocColor = types.StringValue(ccrr.Get("service.tloc.color.value").String())
				}
				if t := ccrr.Get("service.tloc.encap.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocEncapsulation = types.StringValue(ccrr.Get("service.tloc.encap.value").String())
				}
				if t := ccrr.Get("service.tlocList.refId.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceTlocListId = types.StringValue(ccrr.Get("service.tlocList.refId.value").String())
				}
				if t := ccrr.Get("serviceChain.type.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainType = types.StringValue(ccrr.Get("serviceChain.type.value").String())
				}
				if t := ccrr.Get("serviceChain.vpn.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainVpn = types.Int64Value(ccrr.Get("serviceChain.vpn.value").Int())
				}
				if t := ccrr.Get("serviceChain.tloc.ip.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocIp = types.StringValue(ccrr.Get("serviceChain.tloc.ip.value").String())
				}
				if t := ccrr.Get("serviceChain.tloc.color.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocColor = types.StringValue(ccrr.Get("serviceChain.tloc.color.value").String())
				}
				if t := ccrr.Get("serviceChain.tloc.encap.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocEncapsulation = types.StringValue(ccrr.Get("serviceChain.tloc.encap.value").String())
				}
				if t := ccrr.Get("serviceChain.tlocList.refId.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].ServiceChainTlocListId = types.StringValue(ccrr.Get("serviceChain.tlocList.refId.value").String())
				}
				if t := ccrr.Get("tloc.ip.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocIp = types.StringValue(ccrr.Get("tloc.ip.value").String())
				}
				if t := ccrr.Get("tloc.color.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocColor = types.StringValue(ccrr.Get("tloc.color.value").String())
				}
				if t := ccrr.Get("tloc.encap.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocEncapsulation = types.StringValue(ccrr.Get("tloc.encap.value").String())
				}
				if t := ccrr.Get("tlocList.refId.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocListId = types.StringValue(ccrr.Get("tlocList.refId.value").String())
				}
				if t := ccrr.Get("tlocAction.optionType"); t.Exists() && t.String() == "global" {
					data.Sequences[i].ActionEntries[ci].SetParameters[cci].TlocAction = types.StringValue(ccrr.Get("tlocAction.value").String())
				}
			}
		}
	}
}
