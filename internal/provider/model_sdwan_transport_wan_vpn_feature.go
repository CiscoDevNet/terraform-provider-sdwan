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
type TransportWANVPN struct {
	Id                              types.String                      `tfsdk:"id"`
	Version                         types.Int64                       `tfsdk:"version"`
	Name                            types.String                      `tfsdk:"name"`
	Description                     types.String                      `tfsdk:"description"`
	FeatureProfileId                types.String                      `tfsdk:"feature_profile_id"`
	Vpn                             types.Int64                       `tfsdk:"vpn"`
	EnhanceEcmpKeying               types.Bool                        `tfsdk:"enhance_ecmp_keying"`
	EnhanceEcmpKeyingVariable       types.String                      `tfsdk:"enhance_ecmp_keying_variable"`
	PrimaryDnsAddressIpv4           types.String                      `tfsdk:"primary_dns_address_ipv4"`
	PrimaryDnsAddressIpv4Variable   types.String                      `tfsdk:"primary_dns_address_ipv4_variable"`
	SecondaryDnsAddressIpv4         types.String                      `tfsdk:"secondary_dns_address_ipv4"`
	SecondaryDnsAddressIpv4Variable types.String                      `tfsdk:"secondary_dns_address_ipv4_variable"`
	PrimaryDnsAddressIpv6           types.String                      `tfsdk:"primary_dns_address_ipv6"`
	PrimaryDnsAddressIpv6Variable   types.String                      `tfsdk:"primary_dns_address_ipv6_variable"`
	SecondaryDnsAddressIpv6         types.String                      `tfsdk:"secondary_dns_address_ipv6"`
	SecondaryDnsAddressIpv6Variable types.String                      `tfsdk:"secondary_dns_address_ipv6_variable"`
	NewHostMappings                 []TransportWANVPNNewHostMappings  `tfsdk:"new_host_mappings"`
	Ipv4StaticRoutes                []TransportWANVPNIpv4StaticRoutes `tfsdk:"ipv4_static_routes"`
	Ipv6StaticRoutes                []TransportWANVPNIpv6StaticRoutes `tfsdk:"ipv6_static_routes"`
	Services                        []TransportWANVPNServices         `tfsdk:"services"`
	Nat64V4Pools                    []TransportWANVPNNat64V4Pools     `tfsdk:"nat_64_v4_pools"`
}

type TransportWANVPNNewHostMappings struct {
	HostName                  types.String `tfsdk:"host_name"`
	HostNameVariable          types.String `tfsdk:"host_name_variable"`
	ListOfIpAddresses         types.Set    `tfsdk:"list_of_ip_addresses"`
	ListOfIpAddressesVariable types.String `tfsdk:"list_of_ip_addresses_variable"`
}

type TransportWANVPNIpv4StaticRoutes struct {
	NetworkAddress                 types.String                              `tfsdk:"network_address"`
	NetworkAddressVariable         types.String                              `tfsdk:"network_address_variable"`
	SubnetMask                     types.String                              `tfsdk:"subnet_mask"`
	SubnetMaskVariable             types.String                              `tfsdk:"subnet_mask_variable"`
	Gateway                        types.String                              `tfsdk:"gateway"`
	NextHops                       []TransportWANVPNIpv4StaticRoutesNextHops `tfsdk:"next_hops"`
	AdministrativeDistance         types.Int64                               `tfsdk:"administrative_distance"`
	AdministrativeDistanceVariable types.String                              `tfsdk:"administrative_distance_variable"`
}

type TransportWANVPNIpv6StaticRoutes struct {
	Prefix         types.String                              `tfsdk:"prefix"`
	PrefixVariable types.String                              `tfsdk:"prefix_variable"`
	Gateway        types.String                              `tfsdk:"gateway"`
	NextHops       []TransportWANVPNIpv6StaticRoutesNextHops `tfsdk:"next_hops"`
	Null0          types.Bool                                `tfsdk:"null0"`
	Nat            types.String                              `tfsdk:"nat"`
	NatVariable    types.String                              `tfsdk:"nat_variable"`
}

type TransportWANVPNServices struct {
	ServiceType types.String `tfsdk:"service_type"`
}

type TransportWANVPNNat64V4Pools struct {
	Nat64V4PoolName               types.String `tfsdk:"nat64_v4_pool_name"`
	Nat64V4PoolNameVariable       types.String `tfsdk:"nat64_v4_pool_name_variable"`
	Nat64V4PoolRangeStart         types.String `tfsdk:"nat64_v4_pool_range_start"`
	Nat64V4PoolRangeStartVariable types.String `tfsdk:"nat64_v4_pool_range_start_variable"`
	Nat64V4PoolRangeEnd           types.String `tfsdk:"nat64_v4_pool_range_end"`
	Nat64V4PoolRangeEndVariable   types.String `tfsdk:"nat64_v4_pool_range_end_variable"`
	Nat64V4PoolOverload           types.Bool   `tfsdk:"nat64_v4_pool_overload"`
	Nat64V4PoolOverloadVariable   types.String `tfsdk:"nat64_v4_pool_overload_variable"`
}

type TransportWANVPNIpv4StaticRoutesNextHops struct {
	Address                        types.String `tfsdk:"address"`
	AddressVariable                types.String `tfsdk:"address_variable"`
	AdministrativeDistance         types.Int64  `tfsdk:"administrative_distance"`
	AdministrativeDistanceVariable types.String `tfsdk:"administrative_distance_variable"`
}

type TransportWANVPNIpv6StaticRoutesNextHops struct {
	Address                        types.String `tfsdk:"address"`
	AddressVariable                types.String `tfsdk:"address_variable"`
	AdministrativeDistance         types.Int64  `tfsdk:"administrative_distance"`
	AdministrativeDistanceVariable types.String `tfsdk:"administrative_distance_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportWANVPN) getModel() string {
	return "transport_wan_vpn"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportWANVPN) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/wan/vpn", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportWANVPN) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if data.Vpn.IsNull() {
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

	if !data.EnhanceEcmpKeyingVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enhanceEcmpKeying.optionType", "variable")
			body, _ = sjson.Set(body, path+"enhanceEcmpKeying.value", data.EnhanceEcmpKeyingVariable.ValueString())
		}
	} else if data.EnhanceEcmpKeying.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enhanceEcmpKeying.optionType", "default")
			body, _ = sjson.Set(body, path+"enhanceEcmpKeying.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"enhanceEcmpKeying.optionType", "global")
			body, _ = sjson.Set(body, path+"enhanceEcmpKeying.value", data.EnhanceEcmpKeying.ValueBool())
		}
	}

	if !data.PrimaryDnsAddressIpv4Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv4.primaryDnsAddressIpv4.optionType", "variable")
			body, _ = sjson.Set(body, path+"dnsIpv4.primaryDnsAddressIpv4.value", data.PrimaryDnsAddressIpv4Variable.ValueString())
		}
	} else if !data.PrimaryDnsAddressIpv4.IsNull() {
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
	} else if !data.PrimaryDnsAddressIpv6.IsNull() {
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
		for _, item := range data.NewHostMappings {
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

			if !item.ListOfIpAddressesVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "listOfIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "listOfIp.value", item.ListOfIpAddressesVariable.ValueString())
				}
			} else if !item.ListOfIpAddresses.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "listOfIp.optionType", "global")
					var values []string
					item.ListOfIpAddresses.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "listOfIp.value", values)
				}
			}
			body, _ = sjson.SetRaw(body, path+"newHostMapping.-1", itemBody)
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
			if item.Gateway.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "gateway.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "gateway.value", "nextHop")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "gateway.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "gateway.value", item.Gateway.ValueString())
				}
			}
			if true && item.Gateway.ValueString() == "nextHop" {
				itemBody, _ = sjson.Set(itemBody, "nextHop", []interface{}{})
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
					itemBody, _ = sjson.SetRaw(itemBody, "nextHop.-1", itemChildBody)
				}
			}

			if !item.AdministrativeDistanceVariable.IsNull() {
				if true && item.Gateway.ValueString() == "null0" {
					itemBody, _ = sjson.Set(itemBody, "distance.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "distance.value", item.AdministrativeDistanceVariable.ValueString())
				}
			} else if item.AdministrativeDistance.IsNull() {
				if true && item.Gateway.ValueString() == "null0" {
					itemBody, _ = sjson.Set(itemBody, "distance.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "distance.value", 1)
				}
			} else {
				if true && item.Gateway.ValueString() == "null0" {
					itemBody, _ = sjson.Set(itemBody, "distance.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "distance.value", item.AdministrativeDistance.ValueInt64())
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
			body, _ = sjson.SetRaw(body, path+"ipv6Route.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"service", []interface{}{})
		for _, item := range data.Services {
			itemBody := ""
			if !item.ServiceType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "serviceType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "serviceType.value", item.ServiceType.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"service.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"nat64V4Pool", []interface{}{})
		for _, item := range data.Nat64V4Pools {
			itemBody := ""

			if !item.Nat64V4PoolNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.value", item.Nat64V4PoolNameVariable.ValueString())
				}
			} else if !item.Nat64V4PoolName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.value", item.Nat64V4PoolName.ValueString())
				}
			}

			if !item.Nat64V4PoolRangeStartVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.value", item.Nat64V4PoolRangeStartVariable.ValueString())
				}
			} else if !item.Nat64V4PoolRangeStart.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.value", item.Nat64V4PoolRangeStart.ValueString())
				}
			}

			if !item.Nat64V4PoolRangeEndVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.value", item.Nat64V4PoolRangeEndVariable.ValueString())
				}
			} else if !item.Nat64V4PoolRangeEnd.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.value", item.Nat64V4PoolRangeEnd.ValueString())
				}
			}

			if !item.Nat64V4PoolOverloadVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.value", item.Nat64V4PoolOverloadVariable.ValueString())
				}
			} else if item.Nat64V4PoolOverload.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.value", item.Nat64V4PoolOverload.ValueBool())
				}
			}
			body, _ = sjson.SetRaw(body, path+"nat64V4Pool.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportWANVPN) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Vpn = types.Int64Null()

	if t := res.Get(path + "vpnId.optionType"); t.Exists() {
		va := res.Get(path + "vpnId.value")
		if t.String() == "global" {
			data.Vpn = types.Int64Value(va.Int())
		}
	}
	data.EnhanceEcmpKeying = types.BoolNull()
	data.EnhanceEcmpKeyingVariable = types.StringNull()
	if t := res.Get(path + "enhanceEcmpKeying.optionType"); t.Exists() {
		va := res.Get(path + "enhanceEcmpKeying.value")
		if t.String() == "variable" {
			data.EnhanceEcmpKeyingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnhanceEcmpKeying = types.BoolValue(va.Bool())
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
	oldNewHostMappings := data.NewHostMappings
	if value := res.Get(path + "newHostMapping"); value.Exists() && len(value.Array()) > 0 {
		data.NewHostMappings = make([]TransportWANVPNNewHostMappings, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNNewHostMappings{}
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
			item.ListOfIpAddresses = types.SetNull(types.StringType)
			item.ListOfIpAddressesVariable = types.StringNull()
			if t := v.Get("listOfIp.optionType"); t.Exists() {
				va := v.Get("listOfIp.value")
				if t.String() == "variable" {
					item.ListOfIpAddressesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ListOfIpAddresses = helpers.GetStringSet(va.Array())
				}
			}
			data.NewHostMappings = append(data.NewHostMappings, item)
			return true
		})
	} else {
		data.NewHostMappings = nil
	}
	if !fullRead {
		resultNewHostMappings := make([]TransportWANVPNNewHostMappings, 0, len(data.NewHostMappings))
		matchedNewHostMappings := make([]bool, len(data.NewHostMappings))
		for _, oldItem := range oldNewHostMappings {
			for ni := range data.NewHostMappings {
				if matchedNewHostMappings[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.HostNameVariable.ValueString() != "" || data.NewHostMappings[ni].HostNameVariable.ValueString() != "") {
					if oldItem.HostNameVariable.ValueString() != data.NewHostMappings[ni].HostNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.HostName.ValueString() != data.NewHostMappings[ni].HostName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedNewHostMappings[ni] = true
					resultNewHostMappings = append(resultNewHostMappings, data.NewHostMappings[ni])
					break
				}
			}
		}
		for ni := range data.NewHostMappings {
			if !matchedNewHostMappings[ni] {
				resultNewHostMappings = append(resultNewHostMappings, data.NewHostMappings[ni])
			}
		}
		data.NewHostMappings = resultNewHostMappings
	}
	oldIpv4StaticRoutes := data.Ipv4StaticRoutes
	if value := res.Get(path + "ipv4Route"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4StaticRoutes = make([]TransportWANVPNIpv4StaticRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNIpv4StaticRoutes{}
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
			item.Gateway = types.StringNull()

			if t := v.Get("gateway.optionType"); t.Exists() {
				va := v.Get("gateway.value")
				if t.String() == "global" {
					item.Gateway = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("nextHop"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.NextHops = make([]TransportWANVPNIpv4StaticRoutesNextHops, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TransportWANVPNIpv4StaticRoutesNextHops{}
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
			item.AdministrativeDistance = types.Int64Null()
			item.AdministrativeDistanceVariable = types.StringNull()
			if t := v.Get("distance.optionType"); t.Exists() {
				va := v.Get("distance.value")
				if t.String() == "variable" {
					item.AdministrativeDistanceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AdministrativeDistance = types.Int64Value(va.Int())
				}
			}
			data.Ipv4StaticRoutes = append(data.Ipv4StaticRoutes, item)
			return true
		})
	} else {
		data.Ipv4StaticRoutes = nil
	}
	if !fullRead {
		resultIpv4StaticRoutes := make([]TransportWANVPNIpv4StaticRoutes, 0, len(data.Ipv4StaticRoutes))
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
				if keyMatch && (oldItem.SubnetMaskVariable.ValueString() != "" || data.Ipv4StaticRoutes[ni].SubnetMaskVariable.ValueString() != "") {
					if oldItem.SubnetMaskVariable.ValueString() != data.Ipv4StaticRoutes[ni].SubnetMaskVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.SubnetMask.ValueString() != data.Ipv4StaticRoutes[ni].SubnetMask.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.Gateway.ValueString() != data.Ipv4StaticRoutes[ni].Gateway.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedIpv4StaticRoutes[ni] = true
					{
						resultC := make([]TransportWANVPNIpv4StaticRoutesNextHops, 0, len(data.Ipv4StaticRoutes[ni].NextHops))
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
		data.Ipv6StaticRoutes = make([]TransportWANVPNIpv6StaticRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNIpv6StaticRoutes{}
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
				item.NextHops = make([]TransportWANVPNIpv6StaticRoutesNextHops, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TransportWANVPNIpv6StaticRoutesNextHops{}
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
			data.Ipv6StaticRoutes = append(data.Ipv6StaticRoutes, item)
			return true
		})
	} else {
		data.Ipv6StaticRoutes = nil
	}
	if !fullRead {
		resultIpv6StaticRoutes := make([]TransportWANVPNIpv6StaticRoutes, 0, len(data.Ipv6StaticRoutes))
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
						resultC := make([]TransportWANVPNIpv6StaticRoutesNextHops, 0, len(data.Ipv6StaticRoutes[ni].NextHops))
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
		data.Services = make([]TransportWANVPNServices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNServices{}
			item.ServiceType = types.StringNull()

			if t := v.Get("serviceType.optionType"); t.Exists() {
				va := v.Get("serviceType.value")
				if t.String() == "global" {
					item.ServiceType = types.StringValue(va.String())
				}
			}
			data.Services = append(data.Services, item)
			return true
		})
	} else {
		data.Services = nil
	}
	if !fullRead {
		resultServices := make([]TransportWANVPNServices, 0, len(data.Services))
		matchedServices := make([]bool, len(data.Services))
		for _, oldItem := range oldServices {
			for ni := range data.Services {
				if matchedServices[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
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
	oldNat64V4Pools := data.Nat64V4Pools
	if value := res.Get(path + "nat64V4Pool"); value.Exists() && len(value.Array()) > 0 {
		data.Nat64V4Pools = make([]TransportWANVPNNat64V4Pools, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNNat64V4Pools{}
			item.Nat64V4PoolName = types.StringNull()
			item.Nat64V4PoolNameVariable = types.StringNull()
			if t := v.Get("nat64V4PoolName.optionType"); t.Exists() {
				va := v.Get("nat64V4PoolName.value")
				if t.String() == "variable" {
					item.Nat64V4PoolNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Nat64V4PoolName = types.StringValue(va.String())
				}
			}
			item.Nat64V4PoolRangeStart = types.StringNull()
			item.Nat64V4PoolRangeStartVariable = types.StringNull()
			if t := v.Get("nat64V4PoolRangeStart.optionType"); t.Exists() {
				va := v.Get("nat64V4PoolRangeStart.value")
				if t.String() == "variable" {
					item.Nat64V4PoolRangeStartVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Nat64V4PoolRangeStart = types.StringValue(va.String())
				}
			}
			item.Nat64V4PoolRangeEnd = types.StringNull()
			item.Nat64V4PoolRangeEndVariable = types.StringNull()
			if t := v.Get("nat64V4PoolRangeEnd.optionType"); t.Exists() {
				va := v.Get("nat64V4PoolRangeEnd.value")
				if t.String() == "variable" {
					item.Nat64V4PoolRangeEndVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Nat64V4PoolRangeEnd = types.StringValue(va.String())
				}
			}
			item.Nat64V4PoolOverload = types.BoolNull()
			item.Nat64V4PoolOverloadVariable = types.StringNull()
			if t := v.Get("nat64V4PoolOverload.optionType"); t.Exists() {
				va := v.Get("nat64V4PoolOverload.value")
				if t.String() == "variable" {
					item.Nat64V4PoolOverloadVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Nat64V4PoolOverload = types.BoolValue(va.Bool())
				}
			}
			data.Nat64V4Pools = append(data.Nat64V4Pools, item)
			return true
		})
	} else {
		data.Nat64V4Pools = nil
	}
	if !fullRead {
		resultNat64V4Pools := make([]TransportWANVPNNat64V4Pools, 0, len(data.Nat64V4Pools))
		matchedNat64V4Pools := make([]bool, len(data.Nat64V4Pools))
		for _, oldItem := range oldNat64V4Pools {
			for ni := range data.Nat64V4Pools {
				if matchedNat64V4Pools[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.Nat64V4PoolNameVariable.ValueString() != "" || data.Nat64V4Pools[ni].Nat64V4PoolNameVariable.ValueString() != "") {
					if oldItem.Nat64V4PoolNameVariable.ValueString() != data.Nat64V4Pools[ni].Nat64V4PoolNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Nat64V4PoolName.ValueString() != data.Nat64V4Pools[ni].Nat64V4PoolName.ValueString() {
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
}

// End of section. //template:end fromBody
