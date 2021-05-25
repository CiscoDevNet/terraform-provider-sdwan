---
layout: "sdwan"
page_title: "SD-WAN: sdwan_vpn_feature_template"
sidebar_current: "docs-sdwan-data-source-vpn_feature_template"
description: |-
  Data source for SD-WAN VPN Feature Template
---

# sdwan_vpn_feature_template #
Data source for SD-WAN VPN Feature Templates

## Example Usage ##

```hcl

data "sdwan_vpn_feature_template" "example_datasource" {
  template_name = "example_template"
}

```


## Argument Reference ##

* `template_name` - (Required) Unique  VPN Type Feature Template Name

## Attribute Reference ##

* `template_type` - Type of VPN Feature Template
* `template_description` - Long Description of VPN type Feature Template
* `device_type` - Type of device which supports  Feature Template
* `template_min_version` - Minimum Version for the  Feature template
* `template_definition` - Configuration for  VPN type Feature Template, (see [below for nested schema](#nestedblock--template_definition))
* `factory_default` - Boolean flag indicating whether  VPN type Feature Template is factory default or not
* `attached_masters_count` - Number of Device Templates attached to VPN type Feature Template
* `devices_attached` - Number of Devices attached to VPN type Feature Template
* `last_updated_by` - User who updated the VPN type Feature Template latest
* `last_updated_on` - Time for the latest Update of VPN type Feature Template
* `g_template_class` - Template Class type for VPN type Feature Template
* `template_id` - Unique ID for VPN type Feature Template
* `config_type` - Type of configuration for VPN type Feature Template
* `created_on` - Time of creation of VPN type Feature Template
* `rid` - Request id for the VPN type Feature Template
* `feature` - Feature for the  VPN type Feature Template

<a id="nestedblock--template_definition"></a>

## Nested Schema for template_definition
* `vpn_basic` - Basic configuration for the VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_basic))
* `vpn_dns` - DNS configuration for the VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_dns))
* `ipv4_route` - IPv4 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv4_route))
* `ipv6_route` - IPv6 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv6_route))
* `te_service_enabled` - Flag indicating TE service is enabled or not for the VPN type Feature Template
* `service_route` - Service route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--service_route))
* `nat64` - NAT64 for the VPN type Feature Template, (see [below for nested schema](#nestedblock--nat64))


<a id="nestedblock--vpn_basic"></a>

## Nested Schema for vpn_basic
* `vpn_id` - Numeric identifier of the VPN for VPN type of Feature Template
* `name` - Name of VPN for VPN type of Feature Template
* `ecmp_hash_key` - Boolean flag indicating whether ECMP hash key is enabled or not


<a id="nestedblock--vpn_dns"></a>

## Nested Schema for vpn_dns
* `primary_dns_ipv4` - Primary IPv4 address of DNS server for the VPN
* `secondary_dns_ipv4` - Secondary IPv4 address of DNS server for the VPN
* `primary_dns_ipv6` - Primary IPv6 address of DNS server for the VPN
* `secondary_dns_ipv4` - Secondary IPv6 address of DNS server for the VPN
* `vpn_host` - VPN host list for DNS server (see [below for nested schema](#nestedblock--vpn_host))



<a id="nestedblock--ipv4_route"></a>

## Nested Schema for ipv4_route
* `prefix` - IPv4 prefix or address for IPv4 route
* `gateway_type` - Gateway Type for next hop to reach the static route
* `next_hop` - Next hop list (see [below for nested schema](#nestedblock--next_hop)) when gateway_type is `next-hop`
* `null0_distance` - Null0 distance when gateway_type is `null0`


<a id="nestedblock--ipv6_route"></a>

## Nested Schema for ipv6_route
* `prefix` - IPv6 prefix or address for IPv6 route
* `gateway_type` - Gateway Type for next hop to reach the static route
* `next_hop` - Next hop list (see [below for nested schema](#nestedblock--next_hop)) when gateway_type is `next-hop`
* `null0_distance` - Null0 distance when gateway_type is `null0`



<a id="nestedblock--service_route"></a>

## Nested Schema for service_route
* `prefix` - IPv6 or IPv4 prefix or address for service route
* `service_type` - Type of the service


<a id="nestedblock--nat64"></a>

## Nested Schema for nat64
* `pool_name` - NAT64 pool name
* `start_address` - Starting IPv4 address of NAT64 pool
* `end_address` - Ending IPv4 address of NAT64 pool
* `overload` - Flag indicating NAT64 pool overload


<a id="nestedblock--vpn_host"></a>

## Nested Schema for vpn_host
* `hostname` - Host name of the DNS server
* `ip` - List of ip addresses associated with the hostname


<a id="nestedblock--next_hop"></a>

## Nested Schema for next_hop
* `next_hop_address` - IP address of the nexthop
* `next_hop_distance` - Distance of the nexthop
