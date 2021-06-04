---
layout: "sdwan"
page_title: "SD-WAN: sdwan_vpn_feature_template"
sidebar_current: "docs-sdwan-resource-vpn_feature_template"
description: |-
  Manages SD-WAN VPN Type Feature Templates
---

# sdwan_system_feature_template #
Manages SD-WAN VPN Type Feature Templates

## Example Usage ##

```hcl
# example for VPN type of feature template

resource "sdwan_vpn_feature_template" "exmple_resource1" {
  template_name = "example_cisco_vpn"
  template_type = "cisco_vpn"
  factory_default = false
  device_type = [ "vedge-C1111-8P" ]
  template_description = "For testing purpose"
  template_min_version = "15.0.0"
  template_definition {
    vpn_basic {
        vpn_id = 0
    }
  }
}

# example for VPN type of feature template

resource "sdwan_vpn_feature_template" "name1" {
  template_name = "example_vpn_vedge1"
  factory_default = false
  template_type = "vpn-vedge"
  device_type = [ "vedge-cloud" ]
  template_description = "For testing purpose"
  template_min_version = "15.0.0"
  template_definition {
    vpn_basic {
      vpn_id = 0 
      name = "sample_name"
      ecmp_hash_key = "true"
    }
     vpn_dns {
      primary_dns_ipv4 = "1.1.1.1"
      secondary_dns_ipv4 = "1.1.1.2"
      primary_dns_ipv6 = "1::"
      secondary_dns_ipv6 = "2::"
      vpn_host {
        hostname = "h1"
        ip = ["1.1.1.1","2.2.2.2"]
      }
      vpn_host {
        hostname = "h2"
        ip = ["1::"]
      }
     }
    ipv4_route {
      prefix = "1.1.1.1/1"
      gateway_type = "dhcp"
    }
    ipv4_route {
      prefix = "1.1.1.1/2"
      gateway_type = "null0"
      null0_distance = "6"
    }
    ipv4_route {
      prefix = "1.1.1.1/3"
      gateway_type = "null0"
    }
    ipv4_route {
      prefix = "1.1.1.1/4"
      gateway_type = "next-hop"
      next_hop {
        next_hop_address = "1.1.1.1"
      }
      next_hop {
        next_hop_address="2.2.2.2"
        next_hop_distance = "2"
      }
    }
    service_route {
      prefix = "1.1.1.1/4"
      service_type= "FW"
    }
    nat64 {
      pool_name = "smaple_name"
      start_address = "1.1.1.1"
      end_address = "2.2.2.2"
      overload = "true"
    }
    ipv6_route {
      prefix = "1::/1"
      gateway_type = "null0"
      null0_distance = "10"
    }
    ipv6_route {
      prefix = "2::/1"
      gateway_type = "next-hop"
      next_hop {
        next_hop_address = "1::"
        next_hop_distance = "12"
      }
      next_hop {
        next_hop_address = "2::"
      }
    }
    te_service_enabled = true
  }
}
```


## Argument Reference ##

* `template_name` - (Required) Unique VPN Type Feature Template Name, Must not include <, >, !, &, ", or white space; maximum 128 characters
* `template_description` - (Required) Long Description of  System type Feature Template, Should not be empty
* `template_type` - (Required) Template type for VPN type of Feature template, allowed values "vpn-vedge", "cisco_vpn"
* `device_type` - (Required) Type of device which supports  Feature Template, allowed values for template_type `vpn-vedge` "vedge-1000", "vedge-2000", "vedge-cloud", "vedge-5000", "vedge-ISR1100-6G", "vedge-100-B", "vedge-ISR1100-4G", "vedge-100", "vedge-ISR1100-4GLTE", "vedge-100-WM", "vbond", "vedge-100-M", "vedge-ISR1100X-6G", "vedge-ISR1100X-4G", allowed values for template_type `cisco_vpn` "vedge-C8500-12X4QC", "vedge-C1111-4PLTEEA", "vedge-C1161-8P", "vedge-C1117-4PLTEEAW", "vedge-C1121X-8P", "vedge-C8200-1N-4T", "vedge-ISR-4331", "vedge-ISRv", "vedge-C1127X-8PMLTEP", "vedge-C1117-4PMLTEEAWE", "vedge-ISR-4451-X", "vedge-C1113-8PLTEEA", "vedge-IR-1821", "vedge-ASR-1001-X", "vedge-ISR-4321", "vedge-C1116-4PLTEEAWE", "vedge-C1109-4PLTE2P", "vedge-C1121-8P", "vedge-ASR-1002-HX", "vedge-C1111-8PLTEEAW", "vedge-C1112-8PWE", "vedge-C1101-4PLTEP", "vedge-ISR1100-4GLTENA-XE", "vedge-C1111-8PLTELA", "vedge-IR-1835", "vedge-C1121X-8PLTEP", "vedge-IR-1833", "vedge-C8300-1N1S-4T2X", "vedge-C1121-4P", "vedge-ISR-4351", "vedge-C1117-4PLTELA", "vedge-C1116-4PWE", "vedge-nfvis-C8200-UCPE", "vedge-C1113-8PM", "vedge-IR-1831", "vedge-C1127-8PLTEP", "vedge-C1121-8PLTEPW", "vedge-C1113-8PW", "vedge-ASR-1001-HX", "vedge-C1128-8PLTEP", "vedge-C1113-8PLTEEAW", "vedge-C1117-4PW",  "vedge-C1116-4P", "vedge-C1113-8PMLTEEA", "vedge-C1112-8P", "vedge-ISR-4461", "vedge-C1116-4PLTEEA", "vedge-ISR-4221", "vedge-C1117-4PM", "vedge-C1117-4PM", "vedge-C1113-8PLTELAWZ", "vedge-C1117-4PMWE", "vedge-C1109-2PLTEVZ", "vedge-C1113-8P", "vedge-C1117-4P", "vedge-nfvis-ENCS5400", "vedge-C8300-2N2S-6T", "vedge-C1127-8PMLTEP", "vedge-ESR-6300", "vedge-ISR-4221X"  ,"vedge-ISR1100-4GLTEGB-XE vedge-C8500-12X", "vedge-C1109-2PLTEGB", "vedge-CSR-1000v", "vedge-C1113-8PLTEW", "vedge-C1121X-8PLTEPW", "vedge-ISR1100-6G-XE", "vedge-C1121-4PLTEP", "vedge-C1111-8PLTEEA", "vedge-C1117-4PLTEEA", "vedge-C1127X-8PLTEP", "vedge-C1109-2PLTEUS", "vedge-C1112-8PLTEEAWE", "vedge-C1161X-8P", "vedge-C8500L-8S4X", "vedge-C1111-8PW",  "vedge-C1161X-8PLTEP", "vedge-C1101-4PLTEPW", "vedge-ISR1100X-4G-XE", "vedge-IR-1101", "vedge-C1111-4P", "vedge-C1111-4PW", "vedge-C1111-8P", "vedge-C1117-4PMLTEEA", "vedge-C1113-8PLTELA", "vedge-C1111-8PLTELAW", "vedge-C1161-8PLTEP", "vedge-ISR1100X-6G-XE", "vedge-ISR-4431", "vedge-C1101-4P", "vedge-C1109-4PLTE2PW", "vedge-C1113-8PMWE", "vedge-C1118-8P", "vedge-C1126-8PLTEP", "vedge-C8300-1N1S-6T", "vedge-C1121-8PLTEP", "vedge-C8300-2N2S-4T2X", "vedge-C1112-8PLTEEA", "vedge-C1111-4PLTELA", "vedge-ASR-1002-X", "vedge-C1111X-8P", "vedge-C1126X-8PLTEP", "vedge-ASR-1006-X", "vedge-C8000V", "vedge-ISR1100-4G-XE", "vedge-C1117-4PLTELAWZ"  
* `template_min_version` - (Required) Minimum Version for the  Feature template
* `template_definition` - (Required) Configuration for  VPN type Feature Template, (see [below for nested schema](#nestedblock--template_definition))
* `factory_default` - (Required) Boolean flag indicating whether VPN type Feature Template is factory default or not


## Attribute Reference ##

* `attached_masters_count` - Number of Device Templates attached to  VPN type Feature Template
* `devices_attached` - Number of Devices attached to VPN type Feature Template
* `last_updated_by` - User who updated the VPN type Feature Template latest
* `last_updated_on` - Time for the latest Update of VPN type Feature Template
* `g_template_class` - Template Class type for VPN type Feature Template
* `template_id` - Unique ID for VPN type Feature Template
* `config_type` - Type of configuration for VPN type Feature Template
* `created_on` - Time of creation of VPN type Feature Template
* `rid` - Request id for the VPN type Feature Template
* `feature` - Feature for the VPN type Feature Template


<a id="nestedblock--template_definition"></a>

## Nested Schema for template_definition
* `vpn_basic` - (Required) Basic configuration for the  VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_basic))
* `vpn_dns` - (Optional) DNS configuration for the VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_dns))
* `ipv4_route` - (Optional) IPv4 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv4_route))
* `ipv6_route` - (Optional) IPv6 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv6_route))
* `te_service_enabled` - (Optional) Flag indicating TE service is enabled or not for the VPN type Feature Template, can be enabled when `vpn_id` is 0, allowed values: `true`, `false`, default value: `false`
* `service_route` - (Optional) Service route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--service_route))
* `nat64` - (Optional) NAT64 for the VPN type Feature Template, (see [below for nested schema](#nestedblock--nat64)), can be configured for `template_type` "vpn-vedge"


<a id="nestedblock--vpn_basic"></a>

## Nested Schema for vpn_basic
* `vpn_id` - (Required) Numeric identifier of the VPN for VPN type of Feature Template, allowed values: 1, 512
* `name` - (Optional) Name of VPN for VPN type of Feature Template
* `ecmp_hash_key` - (Optional) Boolean flag indicating whether ECMP hash key is enabled or not, allowed values: `true`, `false`, default value: `false`


<a id="nestedblock--vpn_dns"></a>

## Nested Schema for vpn_dns
* `primary_dns_ipv4` - (Optional) Primary IPv4 address of DNS server for the VPN
* `secondary_dns_ipv4` - (Optional) Secondary IPv4 address of DNS server for the VPN, can be set if `primary_dns_ipv4` is set
* `primary_dns_ipv6` - (Optional) Primary IPv6 address of DNS server for the VPN
* `secondary_dns_ipv6` - (Optional) Secondary IPv6 address of DNS server for the VPN, can be set if `primary_dns_ipv6` is set
* `vpn_host` - (Optional) VPN host list for DNS server (see [below for nested schema](#nestedblock--vpn_host))


<a id="nestedblock--ipv4_route"></a>

## Nested Schema for ipv4_route
* `prefix` - (Required) IPv4 prefix or address for IPv4 route
* `gateway_type` - (Required) Gateway Type for next hop to reach the static route, allowed values: "next-hop", "dhcp", "null0"
* `next_hop` - (Optional) Next hop list (see [below for nested schema](#nestedblock--next_hopv4)) when gateway_type is `next-hop`
* `null0_distance` - (Optional) Null0 distance when gateway_type is `null0`, range: [1 - 255]


<a id="nestedblock--ipv6_route"></a>

## Nested Schema for ipv6_route
* `prefix` - (Required) IPv6 prefix or address for IPv6 route
* `gateway_type` - (Required) Gateway Type for next hop to reach the static route, allowed values: "next-hop", "null0"
* `next_hop` - (Optional) Next hop list (see [below for nested schema](#nestedblock--next_hopv6)) when gateway_type is `next-hop`
* `null0_distance` - (Optional) Null0 distance when gateway_type is `null0`, cannot be set for `template_type` "cisco_vpn", range: [1 - 255]


<a id="nestedblock--service_route"></a>

## Nested Schema for service_route
* `prefix` - (Required) IPv4 prefix or address for service route
* `service_type` - (Optional) Type of the service, allowed values for `template_type` "vpn-vedge": "sig", "FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "TE", allowed values for `template_type` "cisco_vpn": "FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "TE"


<a id="nestedblock--nat64"></a>

## Nested Schema for nat64
* `pool_name` - (Required) NAT64 pool name
* `start_address` - (Required) Starting IPv4 address of NAT64 pool
* `end_address` - (Required) Ending IPv4 address of NAT64 pool
* `overload` - (Optional) Flag indicating NAT64 pool overload, allowed valus: `true`, `false`, default values: `false`


<a id="nestedblock--vpn_host"></a>

## Nested Schema for vpn_host
* `hostname` - (Required) Host name of the DNS server
* `ip` - (Required) List of ip addresses associated with the hostname, maximum length: 8


<a id="nestedblock--next_hopv4"></a>

## Nested Schema for next_hop
* `next_hop_address` - (Required) IPv4 address of the nexthop
* `next_hop_distance` - (Optional) Distance of the nexthop, range: [1 - 255]


<a id="nestedblock--next_hopv6"></a>

## Nested Schema for next_hop
* `next_hop_address` - (Required) IPv6 address of the nexthop
* `next_hop_distance` - (Optional) Distance of the nexthop, range: [1 - 255]





