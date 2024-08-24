---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_transport_routing_ospfv3_ipv6_profile_parcel Resource - terraform-provider-sdwan"
subcategory: "Profile Parcels"
description: |-
  This resource can manage a Transport Routing OSPFv3 IPv6 profile parcel.
  Minimum SD-WAN Manager version: 20.12.0
---

# sdwan_transport_routing_ospfv3_ipv6_profile_parcel (Resource)

This resource can manage a Transport Routing OSPFv3 IPv6 profile parcel.
  - Minimum SD-WAN Manager version: `20.12.0`

## Example Usage

```terraform
resource "sdwan_transport_routing_ospfv3_ipv6_profile_parcel" "example" {
  name                           = "Example"
  description                    = "My Example"
  feature_profile_id             = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  router_id                      = "1.2.3.4"
  distance                       = 110
  distance_for_external_routes   = 110
  distance_for_inter_area_routes = 110
  distance_for_intra_area_routes = 110
  reference_bandwidth            = 101
  rfc_1583_compatible            = true
  originate                      = false
  always                         = false
  metric                         = 1
  metric_type                    = "type1"
  spf_calculation_deplay         = 200
  initial_hold_time              = 1000
  maximum_hold_time              = 10000
  filter                         = false
  redistributes = [
    {
      protocol = "static"
    }
  ]
  router_lsa_action           = "on-startup"
  router_lsa_on_startu_p_time = 30
  areas = [
    {
      area_number = 1
      area_type   = "stub"
      interfaces = [
        {
          if_name                 = "GigabitEthernet2"
          hello_interval          = 10
          dead_interval           = 40
          lsa_retransmit_interval = 5
          interface_cost          = 10
          ospf_network_type       = "broadcast"
          passive_interface       = false
          auth_type               = "no-auth"
        }
      ]
      ranges = [
        {
          network      = "3002::/96"
          cost         = 1
          no_advertise = false
        }
      ]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `name` (String) The name of the profile parcel

### Optional

- `always` (Boolean) Always advertise default route
- `always_variable` (String) Variable name
- `areas` (Attributes List) Configure OSPFv3 IPv6 area (see [below for nested schema](#nestedatt--areas))
- `description` (String) The description of the profile parcel
- `distance` (Number) Distance
  - Range: `1`-`254`
  - Default value: `110`
- `distance_for_external_routes` (Number) Set distance for external routes
  - Range: `1`-`254`
  - Default value: `110`
- `distance_for_external_routes_variable` (String) Variable name
- `distance_for_inter_area_routes` (Number) Set distance for inter-area routes
  - Range: `1`-`254`
  - Default value: `110`
- `distance_for_inter_area_routes_variable` (String) Variable name
- `distance_for_intra_area_routes` (Number) Set distance for intra-area routes
  - Range: `1`-`254`
  - Default value: `110`
- `distance_for_intra_area_routes_variable` (String) Variable name
- `distance_variable` (String) Variable name
- `filter` (Boolean) Table map filtered or not
  - Default value: `false`
- `filter_variable` (String) Variable name
- `initial_hold_time` (Number) Set initial hold time between consecutive SPF calculations
  - Range: `1`-`600000`
  - Default value: `1000`
- `initial_hold_time_variable` (String) Variable name
- `maximum_hold_time` (Number) Set maximum hold time between consecutive SPF calculations
  - Range: `1`-`600000`
  - Default value: `10000`
- `maximum_hold_time_variable` (String) Variable name
- `metric` (Number) Set metric used to generate default route <0..16777214>
  - Range: `0`-`16777214`
- `metric_type` (String) Set default route metric type
  - Choices: `type1`, `type2`
- `metric_type_variable` (String) Variable name
- `metric_variable` (String) Variable name
- `originate` (Boolean) Distribute default external route into OSPF disabled
- `redistributes` (Attributes List) Redistribute routes (see [below for nested schema](#nestedatt--redistributes))
- `reference_bandwidth` (Number) Set reference bandwidth method to assign OSPF cost
  - Range: `1`-`4294967`
  - Default value: `100`
- `reference_bandwidth_variable` (String) Variable name
- `rfc_1583_compatible` (Boolean) Calculate summary route cost based on RFC 1583
  - Default value: `true`
- `rfc_1583_compatible_variable` (String) Variable name
- `route_policy_id` (String)
- `router_id` (String) Set OSPF router ID to override system IP address
- `router_id_variable` (String) Variable name
- `router_lsa_action` (String) Not advertise maximum metric Router LSA policy by default
- `router_lsa_on_startu_p_time` (Number) Set how long to advertise maximum metric after router boot up
  - Range: `5`-`86400`
- `router_lsa_on_startu_p_time_variable` (String) Variable name
- `spf_calculation_deplay` (Number) Set delay from first change received until performing SPF calculation
  - Range: `1`-`600000`
  - Default value: `200`
- `spf_calculation_deplay_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the profile parcel
- `version` (Number) The version of the profile parcel

<a id="nestedatt--areas"></a>
### Nested Schema for `areas`

Optional:

- `always_translate` (Boolean) Always translate type7 LSAs
- `always_translate_variable` (String) Variable name
- `area_number` (Number) Set OSPF area number
  - Range: `0`-`4294967295`
- `area_number_variable` (String) Variable name
- `area_type` (String) stub area type
  - Choices: `stub`
- `interfaces` (Attributes List) Set OSPF interface parameters (see [below for nested schema](#nestedatt--areas--interfaces))
- `no_summary` (Boolean) Do not inject inter-area routes
- `no_summary_variable` (String) Variable name
- `ranges` (Attributes List) Summarize OSPF routes at an area boundary (see [below for nested schema](#nestedatt--areas--ranges))

<a id="nestedatt--areas--interfaces"></a>
### Nested Schema for `areas.interfaces`

Optional:

- `auth_key` (String) Set OSPF interface authentication IPSEC key
- `auth_key_variable` (String) Variable name
- `auth_type` (String) No Authentication by default
  - Choices: `no-auth`
- `dead_interval` (Number) Set interval after which neighbor is declared to be down
  - Range: `1`-`65535`
  - Default value: `40`
- `dead_interval_variable` (String) Variable name
- `hello_interval` (Number) Set interval between OSPF hello packets
  - Range: `1`-`65535`
  - Default value: `10`
- `hello_interval_variable` (String) Variable name
- `if_name` (String) Set interface name
- `if_name_variable` (String) Variable name
- `interface_cost` (Number) Set cost of OSPF interface
  - Range: `1`-`65535`
- `interface_cost_variable` (String) Variable name
- `lsa_retransmit_interval` (Number) Set time between retransmitting LSAs
  - Range: `1`-`65535`
  - Default value: `5`
- `lsa_retransmit_interval_variable` (String) Variable name
- `ospf_network_type` (String) Set the OSPF network type
  - Choices: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`
- `ospf_network_type_variable` (String) Variable name
- `passive_interface` (Boolean) Set the interface to advertise its address, but not to actively run OSPF
  - Default value: `false`
- `passive_interface_variable` (String) Variable name
- `spi` (Number) Set OSPF interface authentication IPSec SPI, range 256..4294967295
  - Range: `256`-`4294967295`
- `spi_variable` (String) Variable name


<a id="nestedatt--areas--ranges"></a>
### Nested Schema for `areas.ranges`

Optional:

- `cost` (Number) Set cost for this range
  - Range: `0`-`16777214`
- `cost_variable` (String) Variable name
- `network` (String) IPv6 prefix,for example 2001::/64
- `network_variable` (String) Variable name
- `no_advertise` (Boolean) Do not advertise this range
  - Default value: `false`
- `no_advertise_variable` (String) Variable name



<a id="nestedatt--redistributes"></a>
### Nested Schema for `redistributes`

Optional:

- `protocol` (String) Set the protocol
  - Choices: `connected`, `static`, `omp`, `bgp`, `eigrp`
- `protocol_variable` (String) Variable name
- `route_policy_id` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_transport_routing_ospfv3_ipv6_profile_parcel.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```