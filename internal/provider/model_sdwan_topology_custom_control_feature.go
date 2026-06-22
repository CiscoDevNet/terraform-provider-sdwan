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
	if data.TargetInboundSites.IsNull() {
		if true && data.TargetLevel.ValueString() == "SITE" {
			body, _ = sjson.Set(body, path+"target.inboundSites.optionType", "global")
			body, _ = sjson.Set(body, path+"target.inboundSites.value", []interface{}{})
		}
	} else {
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
							itemChildBody, _ = sjson.Set(itemChildBody, "colorList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "colorList.refId.value", childItem.ColorListId.ValueString())
						}
					}
					if !childItem.CommunityListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "community.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "community.refId.value", childItem.CommunityListId.ValueString())
						}
					}
					if !childItem.ExpandedCommunityListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "expandedCommunity.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "expandedCommunity.refId.value", childItem.ExpandedCommunityListId.ValueString())
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
							itemChildBody, _ = sjson.Set(itemChildBody, "tlocList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "tlocList.refId.value", childItem.TlocListId.ValueString())
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
							itemChildBody, _ = sjson.Set(itemChildBody, "prefixList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefixList.refId.value", childItem.PrefixListId.ValueString())
						}
					}
					if !childItem.Ipv6PrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6prefixList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6prefixList.refId.value", childItem.Ipv6PrefixListId.ValueString())
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
func (data *TopologyCustomControl) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
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
	oldTargetInboundRegions := data.TargetInboundRegions
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
	} else {
		data.TargetInboundRegions = nil
	}
	if !fullRead && data.TargetInboundRegions != nil {
		resultTargetInboundRegions := make([]TopologyCustomControlTargetInboundRegions, 0, len(data.TargetInboundRegions))
		matchedTargetInboundRegions := make([]bool, len(data.TargetInboundRegions))
		for _, oldItem := range oldTargetInboundRegions {
			for ni := range data.TargetInboundRegions {
				if matchedTargetInboundRegions[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Region.ValueString() != data.TargetInboundRegions[ni].Region.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SubRegions).ValueString() != helpers.GetStringFromSet(data.TargetInboundRegions[ni].SubRegions).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedTargetInboundRegions[ni] = true
					resultTargetInboundRegions = append(resultTargetInboundRegions, data.TargetInboundRegions[ni])
					break
				}
			}
		}
		for ni := range data.TargetInboundRegions {
			if !matchedTargetInboundRegions[ni] {
				resultTargetInboundRegions = append(resultTargetInboundRegions, data.TargetInboundRegions[ni])
			}
		}
		data.TargetInboundRegions = resultTargetInboundRegions
	}
	oldTargetOutboundRegions := data.TargetOutboundRegions
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
	} else {
		data.TargetOutboundRegions = nil
	}
	if !fullRead && data.TargetOutboundRegions != nil {
		resultTargetOutboundRegions := make([]TopologyCustomControlTargetOutboundRegions, 0, len(data.TargetOutboundRegions))
		matchedTargetOutboundRegions := make([]bool, len(data.TargetOutboundRegions))
		for _, oldItem := range oldTargetOutboundRegions {
			for ni := range data.TargetOutboundRegions {
				if matchedTargetOutboundRegions[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Region.ValueString() != data.TargetOutboundRegions[ni].Region.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SubRegions).ValueString() != helpers.GetStringFromSet(data.TargetOutboundRegions[ni].SubRegions).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedTargetOutboundRegions[ni] = true
					resultTargetOutboundRegions = append(resultTargetOutboundRegions, data.TargetOutboundRegions[ni])
					break
				}
			}
		}
		for ni := range data.TargetOutboundRegions {
			if !matchedTargetOutboundRegions[ni] {
				resultTargetOutboundRegions = append(resultTargetOutboundRegions, data.TargetOutboundRegions[ni])
			}
		}
		data.TargetOutboundRegions = resultTargetOutboundRegions
	}
	oldSequences := data.Sequences
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

					if t := cv.Get("colorList.refId.optionType"); t.Exists() {
						va := cv.Get("colorList.refId.value")
						if t.String() == "global" {
							cItem.ColorListId = types.StringValue(va.String())
						}
					}
					cItem.CommunityListId = types.StringNull()

					if t := cv.Get("community.refId.optionType"); t.Exists() {
						va := cv.Get("community.refId.value")
						if t.String() == "global" {
							cItem.CommunityListId = types.StringValue(va.String())
						}
					}
					cItem.ExpandedCommunityListId = types.StringNull()

					if t := cv.Get("expandedCommunity.refId.optionType"); t.Exists() {
						va := cv.Get("expandedCommunity.refId.value")
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

					if t := cv.Get("tlocList.refId.optionType"); t.Exists() {
						va := cv.Get("tlocList.refId.value")
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

					if t := cv.Get("prefixList.refId.optionType"); t.Exists() {
						va := cv.Get("prefixList.refId.value")
						if t.String() == "global" {
							cItem.PrefixListId = types.StringValue(va.String())
						}
					}
					cItem.Ipv6PrefixListId = types.StringNull()

					if t := cv.Get("ipv6prefixList.refId.optionType"); t.Exists() {
						va := cv.Get("ipv6prefixList.refId.value")
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
	} else {
		data.Sequences = nil
	}
	if !fullRead && data.Sequences != nil {
		resultSequences := make([]TopologyCustomControlSequences, 0, len(data.Sequences))
		matchedSequences := make([]bool, len(data.Sequences))
		for _, oldItem := range oldSequences {
			for ni := range data.Sequences {
				if matchedSequences[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Id.ValueInt64() != data.Sequences[ni].Id.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedSequences[ni] = true
					if data.Sequences[ni].MatchEntries != nil {
						resultC := make([]TopologyCustomControlSequencesMatchEntries, 0, len(data.Sequences[ni].MatchEntries))
						matchedC := make([]bool, len(data.Sequences[ni].MatchEntries))
						for _, oldCItem := range oldItem.MatchEntries {
							for nci := range data.Sequences[ni].MatchEntries {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.ColorListId.ValueString() != data.Sequences[ni].MatchEntries[nci].ColorListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.CommunityListId.ValueString() != data.Sequences[ni].MatchEntries[nci].CommunityListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.ExpandedCommunityListId.ValueString() != data.Sequences[ni].MatchEntries[nci].ExpandedCommunityListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.OmpTag.ValueInt64() != data.Sequences[ni].MatchEntries[nci].OmpTag.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Origin.ValueString() != data.Sequences[ni].MatchEntries[nci].Origin.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Originator.ValueString() != data.Sequences[ni].MatchEntries[nci].Originator.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Preference.ValueInt64() != data.Sequences[ni].MatchEntries[nci].Preference.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.Site).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].Site).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Role.ValueString() != data.Sequences[ni].MatchEntries[nci].Role.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.PathType.ValueString() != data.Sequences[ni].MatchEntries[nci].PathType.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.TlocListId.ValueString() != data.Sequences[ni].MatchEntries[nci].TlocListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.Vpn).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].Vpn).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.PrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].PrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Ipv6PrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].Ipv6PrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Carrier.ValueString() != data.Sequences[ni].MatchEntries[nci].Carrier.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.DomainId.ValueInt64() != data.Sequences[ni].MatchEntries[nci].DomainId.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.GroupId.ValueInt64() != data.Sequences[ni].MatchEntries[nci].GroupId.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.TlocIp.ValueString() != data.Sequences[ni].MatchEntries[nci].TlocIp.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.TlocColor.ValueString() != data.Sequences[ni].MatchEntries[nci].TlocColor.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.TlocEncapsulation.ValueString() != data.Sequences[ni].MatchEntries[nci].TlocEncapsulation.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									if data.Sequences[ni].MatchEntries[nci].MatchRegions != nil {
										resultCC := make([]TopologyCustomControlSequencesMatchEntriesMatchRegions, 0, len(data.Sequences[ni].MatchEntries[nci].MatchRegions))
										matchedCC := make([]bool, len(data.Sequences[ni].MatchEntries[nci].MatchRegions))
										for _, oldCCItem := range oldCItem.MatchRegions {
											for ncci := range data.Sequences[ni].MatchEntries[nci].MatchRegions {
												if matchedCC[ncci] {
													continue
												}
												keyMatchCC := true
												if keyMatchCC {
													if oldCCItem.Region.ValueString() != data.Sequences[ni].MatchEntries[nci].MatchRegions[ncci].Region.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if helpers.GetStringFromSet(oldCCItem.SubRegions).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].MatchRegions[ncci].SubRegions).ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													matchedCC[ncci] = true
													resultCC = append(resultCC, data.Sequences[ni].MatchEntries[nci].MatchRegions[ncci])
													break
												}
											}
										}
										for ncci := range data.Sequences[ni].MatchEntries[nci].MatchRegions {
											if !matchedCC[ncci] {
												resultCC = append(resultCC, data.Sequences[ni].MatchEntries[nci].MatchRegions[ncci])
											}
										}
										data.Sequences[ni].MatchEntries[nci].MatchRegions = resultCC
									}
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
					if data.Sequences[ni].ActionEntries != nil {
						resultC := make([]TopologyCustomControlSequencesActionEntries, 0, len(data.Sequences[ni].ActionEntries))
						matchedC := make([]bool, len(data.Sequences[ni].ActionEntries))
						for _, oldCItem := range oldItem.ActionEntries {
							for nci := range data.Sequences[ni].ActionEntries {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.ExportToVpn).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].ActionEntries[nci].ExportToVpn).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									if data.Sequences[ni].ActionEntries[nci].SetParameters != nil {
										resultCC := make([]TopologyCustomControlSequencesActionEntriesSetParameters, 0, len(data.Sequences[ni].ActionEntries[nci].SetParameters))
										matchedCC := make([]bool, len(data.Sequences[ni].ActionEntries[nci].SetParameters))
										for _, oldCCItem := range oldCItem.SetParameters {
											for ncci := range data.Sequences[ni].ActionEntries[nci].SetParameters {
												if matchedCC[ncci] {
													continue
												}
												keyMatchCC := true
												if keyMatchCC {
													if oldCCItem.Preference.ValueInt64() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].Preference.ValueInt64() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.OmpTag.ValueInt64() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].OmpTag.ValueInt64() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.Community.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].Community.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.CommunityAdditive.ValueBool() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].CommunityAdditive.ValueBool() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.Affinity.ValueInt64() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].Affinity.ValueInt64() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceType.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceType.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceVpn.ValueInt64() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceVpn.ValueInt64() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceTlocIp.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceTlocIp.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceTlocColor.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceTlocColor.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceTlocEncapsulation.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceTlocEncapsulation.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceTlocListId.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceTlocListId.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainType.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceChainType.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainVpn.ValueInt64() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceChainVpn.ValueInt64() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainTlocIp.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceChainTlocIp.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainTlocColor.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceChainTlocColor.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainTlocEncapsulation.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceChainTlocEncapsulation.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.ServiceChainTlocListId.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].ServiceChainTlocListId.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.TlocIp.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].TlocIp.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.TlocColor.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].TlocColor.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.TlocEncapsulation.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].TlocEncapsulation.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.TlocListId.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].TlocListId.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													if oldCCItem.TlocAction.ValueString() != data.Sequences[ni].ActionEntries[nci].SetParameters[ncci].TlocAction.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													matchedCC[ncci] = true
													resultCC = append(resultCC, data.Sequences[ni].ActionEntries[nci].SetParameters[ncci])
													break
												}
											}
										}
										for ncci := range data.Sequences[ni].ActionEntries[nci].SetParameters {
											if !matchedCC[ncci] {
												resultCC = append(resultCC, data.Sequences[ni].ActionEntries[nci].SetParameters[ncci])
											}
										}
										data.Sequences[ni].ActionEntries[nci].SetParameters = resultCC
									}
									resultC = append(resultC, data.Sequences[ni].ActionEntries[nci])
									break
								}
							}
						}
						for nci := range data.Sequences[ni].ActionEntries {
							if !matchedC[nci] {
								resultC = append(resultC, data.Sequences[ni].ActionEntries[nci])
							}
						}
						data.Sequences[ni].ActionEntries = resultC
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
