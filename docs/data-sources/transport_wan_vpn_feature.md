---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_transport_wan_vpn_feature Data Source - terraform-provider-sdwan"
subcategory: "Features - Transport"
description: |-
  This data source can read the Transport WAN VPN Feature.
---

# sdwan_transport_wan_vpn_feature (Data Source)

This data source can read the Transport WAN VPN Feature.

## Example Usage

```terraform
data "sdwan_transport_wan_vpn_feature" "example" {
  id                 = "f6b2c44c-693c-4763-b010-895aa3d236bd"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `id` (String) The id of the Feature

### Read-Only

- `description` (String) The description of the Feature
- `enhance_ecmp_keying` (Boolean) Enhance ECMP Keying
- `enhance_ecmp_keying_variable` (String) Variable name
- `ipv4_static_routes` (Attributes List) IPv4 Static Route (see [below for nested schema](#nestedatt--ipv4_static_routes))
- `ipv6_static_routes` (Attributes List) IPv6 Static Route (see [below for nested schema](#nestedatt--ipv6_static_routes))
- `name` (String) The name of the Feature
- `nat_64_v4_pools` (Attributes List) NAT64 V4 Pool (see [below for nested schema](#nestedatt--nat_64_v4_pools))
- `new_host_mappings` (Attributes List) (see [below for nested schema](#nestedatt--new_host_mappings))
- `primary_dns_address_ipv4` (String) Primary DNS Address (IPv4)
- `primary_dns_address_ipv4_variable` (String) Variable name
- `primary_dns_address_ipv6` (String) Primary DNS Address (IPv6)
- `primary_dns_address_ipv6_variable` (String) Variable name
- `secondary_dns_address_ipv4` (String) Secondary DNS Address (IPv4)
- `secondary_dns_address_ipv4_variable` (String) Variable name
- `secondary_dns_address_ipv6` (String) Secondary DNS Address (IPv6)
- `secondary_dns_address_ipv6_variable` (String) Variable name
- `services` (Attributes List) Service (see [below for nested schema](#nestedatt--services))
- `version` (Number) The version of the Feature
- `vpn` (Number) VPN

<a id="nestedatt--ipv4_static_routes"></a>
### Nested Schema for `ipv4_static_routes`

Read-Only:

- `administrative_distance` (Number) Administrative distance
- `administrative_distance_variable` (String) Variable name
- `gateway` (String) Gateway
- `network_address` (String) IP Address
- `network_address_variable` (String) Variable name
- `next_hops` (Attributes List) IPv4 Route Gateway Next Hop (see [below for nested schema](#nestedatt--ipv4_static_routes--next_hops))
- `subnet_mask` (String) Subnet Mask
- `subnet_mask_variable` (String) Variable name

<a id="nestedatt--ipv4_static_routes--next_hops"></a>
### Nested Schema for `ipv4_static_routes.next_hops`

Read-Only:

- `address` (String) Address
- `address_variable` (String) Variable name
- `administrative_distance` (Number) Administrative distance
- `administrative_distance_variable` (String) Variable name



<a id="nestedatt--ipv6_static_routes"></a>
### Nested Schema for `ipv6_static_routes`

Read-Only:

- `gateway` (String) Gateway
- `nat` (String) IPv6 Nat
- `nat_variable` (String) Variable name
- `next_hops` (Attributes List) IPv6 Route Gateway Next Hop (see [below for nested schema](#nestedatt--ipv6_static_routes--next_hops))
- `null0` (Boolean) IPv6 Route Gateway Next Hop
- `prefix` (String) Prefix
- `prefix_variable` (String) Variable name

<a id="nestedatt--ipv6_static_routes--next_hops"></a>
### Nested Schema for `ipv6_static_routes.next_hops`

Read-Only:

- `address` (String) Address
- `address_variable` (String) Variable name
- `administrative_distance` (Number) Administrative distance
- `administrative_distance_variable` (String) Variable name



<a id="nestedatt--nat_64_v4_pools"></a>
### Nested Schema for `nat_64_v4_pools`

Read-Only:

- `nat64_v4_pool_name` (String) NAT64 v4 Pool Name
- `nat64_v4_pool_name_variable` (String) Variable name
- `nat64_v4_pool_overload` (Boolean) NAT64 Overload
- `nat64_v4_pool_overload_variable` (String) Variable name
- `nat64_v4_pool_range_end` (String) NAT64 Pool Range End
- `nat64_v4_pool_range_end_variable` (String) Variable name
- `nat64_v4_pool_range_start` (String) NAT64 Pool Range Start
- `nat64_v4_pool_range_start_variable` (String) Variable name


<a id="nestedatt--new_host_mappings"></a>
### Nested Schema for `new_host_mappings`

Read-Only:

- `host_name` (String) Hostname
- `host_name_variable` (String) Variable name
- `list_of_ip_addresses` (Set of String) List of IP
- `list_of_ip_addresses_variable` (String) Variable name


<a id="nestedatt--services"></a>
### Nested Schema for `services`

Read-Only:

- `service_type` (String) Service Type
