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
					} else if !childItem.AdministrativeDistance.IsNull() {
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
func (data *TransportWANVPN) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "newHostMapping"); value.Exists() {
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
	}
	if value := res.Get(path + "ipv4Route"); value.Exists() {
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
			if cValue := v.Get("nextHop"); cValue.Exists() {
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
	}
	if value := res.Get(path + "ipv6Route"); value.Exists() {
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
			if cValue := v.Get("oneOfIpRoute.nextHopContainer.nextHop"); cValue.Exists() {
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
	}
	if value := res.Get(path + "nat64V4Pool"); value.Exists() {
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
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportWANVPN) updateFromBody(ctx context.Context, res gjson.Result) {
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
	for i := range data.NewHostMappings {
		keys := [...]string{"hostName"}
		keyValues := [...]string{data.NewHostMappings[i].HostName.ValueString()}
		keyValuesVariables := [...]string{data.NewHostMappings[i].HostNameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "newHostMapping").ForEach(
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
		data.NewHostMappings[i].HostName = types.StringNull()
		data.NewHostMappings[i].HostNameVariable = types.StringNull()
		if t := r.Get("hostName.optionType"); t.Exists() {
			va := r.Get("hostName.value")
			if t.String() == "variable" {
				data.NewHostMappings[i].HostNameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NewHostMappings[i].HostName = types.StringValue(va.String())
			}
		}
		data.NewHostMappings[i].ListOfIpAddresses = types.SetNull(types.StringType)
		data.NewHostMappings[i].ListOfIpAddressesVariable = types.StringNull()
		if t := r.Get("listOfIp.optionType"); t.Exists() {
			va := r.Get("listOfIp.value")
			if t.String() == "variable" {
				data.NewHostMappings[i].ListOfIpAddressesVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NewHostMappings[i].ListOfIpAddresses = helpers.GetStringSet(va.Array())
			}
		}
	}
	for i := range data.Ipv4StaticRoutes {
		keys := [...]string{"prefix.ipAddress", "prefix.subnetMask", "gateway"}
		keyValues := [...]string{data.Ipv4StaticRoutes[i].NetworkAddress.ValueString(), data.Ipv4StaticRoutes[i].SubnetMask.ValueString(), data.Ipv4StaticRoutes[i].Gateway.ValueString()}
		keyValuesVariables := [...]string{data.Ipv4StaticRoutes[i].NetworkAddressVariable.ValueString(), data.Ipv4StaticRoutes[i].SubnetMaskVariable.ValueString(), ""}

		var r gjson.Result
		res.Get(path + "ipv4Route").ForEach(
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
		data.Ipv4StaticRoutes[i].Gateway = types.StringNull()

		if t := r.Get("gateway.optionType"); t.Exists() {
			va := r.Get("gateway.value")
			if t.String() == "global" {
				data.Ipv4StaticRoutes[i].Gateway = types.StringValue(va.String())
			}
		}
		for ci := range data.Ipv4StaticRoutes[i].NextHops {
			keys := [...]string{"address"}
			keyValues := [...]string{data.Ipv4StaticRoutes[i].NextHops[ci].Address.ValueString()}
			keyValuesVariables := [...]string{data.Ipv4StaticRoutes[i].NextHops[ci].AddressVariable.ValueString()}

			var cr gjson.Result
			r.Get("nextHop").ForEach(
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
		data.Ipv4StaticRoutes[i].AdministrativeDistance = types.Int64Null()
		data.Ipv4StaticRoutes[i].AdministrativeDistanceVariable = types.StringNull()
		if t := r.Get("distance.optionType"); t.Exists() {
			va := r.Get("distance.value")
			if t.String() == "variable" {
				data.Ipv4StaticRoutes[i].AdministrativeDistanceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4StaticRoutes[i].AdministrativeDistance = types.Int64Value(va.Int())
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
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "service").ForEach(
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
		data.Services[i].ServiceType = types.StringNull()

		if t := r.Get("serviceType.optionType"); t.Exists() {
			va := r.Get("serviceType.value")
			if t.String() == "global" {
				data.Services[i].ServiceType = types.StringValue(va.String())
			}
		}
	}
	for i := range data.Nat64V4Pools {
		keys := [...]string{"nat64V4PoolName"}
		keyValues := [...]string{data.Nat64V4Pools[i].Nat64V4PoolName.ValueString()}
		keyValuesVariables := [...]string{data.Nat64V4Pools[i].Nat64V4PoolNameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "nat64V4Pool").ForEach(
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
		data.Nat64V4Pools[i].Nat64V4PoolName = types.StringNull()
		data.Nat64V4Pools[i].Nat64V4PoolNameVariable = types.StringNull()
		if t := r.Get("nat64V4PoolName.optionType"); t.Exists() {
			va := r.Get("nat64V4PoolName.value")
			if t.String() == "variable" {
				data.Nat64V4Pools[i].Nat64V4PoolNameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Nat64V4Pools[i].Nat64V4PoolName = types.StringValue(va.String())
			}
		}
		data.Nat64V4Pools[i].Nat64V4PoolRangeStart = types.StringNull()
		data.Nat64V4Pools[i].Nat64V4PoolRangeStartVariable = types.StringNull()
		if t := r.Get("nat64V4PoolRangeStart.optionType"); t.Exists() {
			va := r.Get("nat64V4PoolRangeStart.value")
			if t.String() == "variable" {
				data.Nat64V4Pools[i].Nat64V4PoolRangeStartVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Nat64V4Pools[i].Nat64V4PoolRangeStart = types.StringValue(va.String())
			}
		}
		data.Nat64V4Pools[i].Nat64V4PoolRangeEnd = types.StringNull()
		data.Nat64V4Pools[i].Nat64V4PoolRangeEndVariable = types.StringNull()
		if t := r.Get("nat64V4PoolRangeEnd.optionType"); t.Exists() {
			va := r.Get("nat64V4PoolRangeEnd.value")
			if t.String() == "variable" {
				data.Nat64V4Pools[i].Nat64V4PoolRangeEndVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Nat64V4Pools[i].Nat64V4PoolRangeEnd = types.StringValue(va.String())
			}
		}
		data.Nat64V4Pools[i].Nat64V4PoolOverload = types.BoolNull()
		data.Nat64V4Pools[i].Nat64V4PoolOverloadVariable = types.StringNull()
		if t := r.Get("nat64V4PoolOverload.optionType"); t.Exists() {
			va := r.Get("nat64V4PoolOverload.value")
			if t.String() == "variable" {
				data.Nat64V4Pools[i].Nat64V4PoolOverloadVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Nat64V4Pools[i].Nat64V4PoolOverload = types.BoolValue(va.Bool())
			}
		}
	}
}

// End of section. //template:end updateFromBody
