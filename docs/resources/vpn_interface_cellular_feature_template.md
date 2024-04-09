---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_vpn_interface_cellular_feature_template Resource - terraform-provider-sdwan"
subcategory: "Feature Templates"
description: |-
  This resource can manage a VPN Interface Cellular feature template.
    - Minimum SD-WAN Manager version: 15.0.0
---

# sdwan_vpn_interface_cellular_feature_template (Resource)

This resource can manage a VPN Interface Cellular feature template.
  - Minimum SD-WAN Manager version: `15.0.0`

## Example Usage

```terraform
resource "sdwan_vpn_interface_cellular_feature_template" "example" {
  name                    = "Example"
  description             = "My Example"
  device_types            = ["vedge-C8000V"]
  cellular_interface_name = "Cellular1"
  interface_description   = "My Description"
  ipv6_access_lists = [
    {
      direction = "in"
      acl_name  = "ACL1"
    }
  ]
  ipv4_dhcp_helper     = ["6.6.6.6"]
  tracker              = ["tracker1"]
  nat                  = true
  nat_refresh_mode     = "outbound"
  nat_udp_timeout      = 1
  nat_tcp_timeout      = 60
  nat_block_icmp_error = true
  nat_response_to_ping = false
  nat_port_forwards = [
    {
      port_start_range   = 0
      port_end_range     = 65530
      protocol           = "tcp"
      private_vpn        = 65530
      private_ip_address = "1.2.3.4"
    }
  ]
  enable_core_region = true
  core_region        = "core"
  secondary_region   = "off"
  tunnel_interface_encapsulations = [
    {
      encapsulation = "gre"
      preference    = 4294967
      weight        = 250
    }
  ]
  tunnel_interface_groups                        = [42949672]
  tunnel_interface_border                        = true
  per_tunnel_qos                                 = true
  per_tunnel_qos_aggregator                      = false
  tunnel_qos_mode                                = "spoke"
  tunnel_interface_color                         = "custom1"
  tunnel_interface_last_resort_circuit           = false
  tunnel_interface_low_bandwidth_link            = false
  tunnel_interface_tunnel_tcp_mss                = 1460
  tunnel_interface_clear_dont_fragment           = false
  tunnel_interface_network_broadcast             = false
  tunnel_interface_max_control_connections       = 8
  tunnel_interface_control_connections           = true
  tunnel_interface_vbond_as_stun_server          = false
  tunnel_interface_exclude_controller_group_list = [100]
  tunnel_interface_vmanage_connection_preference = 5
  tunnel_interface_port_hop                      = false
  tunnel_interface_color_restrict                = false
  tunnel_interface_carrier                       = "carrier1"
  tunnel_interface_nat_refresh_interval          = 15
  tunnel_interface_hello_interval                = 1000
  tunnel_interface_hello_tolerance               = 12
  tunnel_interface_bind_loopback_tunnel          = "12"
  tunnel_interface_allow_all                     = false
  tunnel_interface_allow_bgp                     = false
  tunnel_interface_allow_dhcp                    = true
  tunnel_interface_allow_dns                     = true
  tunnel_interface_allow_icmp                    = true
  tunnel_interface_allow_ssh                     = false
  tunnel_interface_allow_ntp                     = false
  tunnel_interface_allow_netconf                 = false
  tunnel_interface_allow_ospf                    = false
  tunnel_interface_allow_stun                    = false
  tunnel_interface_allow_snmp                    = false
  tunnel_interface_allow_https                   = true
  clear_dont_fragment_bit                        = false
  pmtu_discovery                                 = false
  ip_mtu                                         = 1500
  static_ingress_qos                             = 6
  tcp_mss                                        = 720
  tloc_extension                                 = "tloc"
  ip_directed_broadcast                          = true
  shutdown                                       = true
  autonegotiate                                  = true
  qos_adaptive_period                            = 15
  qos_adaptive_bandwidth_downstream              = 10000
  qos_adaptive_min_downstream                    = 100
  qos_adaptive_max_downstream                    = 100000
  qos_adaptive_bandwidth_upstream                = 10000
  qos_adaptive_min_upstream                      = 100
  qos_adaptive_max_upstream                      = 100000
  shaping_rate                                   = 10000000
  qos_map                                        = "test"
  qos_map_vpn                                    = "test"
  bandwidth_upstream                             = 214748300
  bandwidth_downstream                           = 214748300
  write_rule                                     = "RULE1"
  ipv4_access_lists = [
    {
      direction = "in"
      acl_name  = "ACL2"
    }
  ]
  policers = [
    {
      direction    = "in"
      policer_name = "example"
    }
  ]
  static_arps = [
    {
      ip_address = "1.2.3.4"
      mac        = "00-B0-D0-63-C2-26"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the feature template
- `device_types` (Set of String) List of supported device types
  - Choices: `vedge-C8000V`, `vedge-C8300-1N1S-4T2X`, `vedge-C8300-1N1S-6T`, `vedge-C8300-2N2S-6T`, `vedge-C8300-2N2S-4T2X`, `vedge-C8500-12X4QC`, `vedge-C8500-12X`, `vedge-C8500-20X6C`, `vedge-C8500L-8S4X`, `vedge-C8200-1N-4T`, `vedge-C8200L-1N-4T`
- `name` (String) The name of the feature template

### Optional

- `autonegotiate` (Boolean) Link autonegotiation
  - Default value: `true`
- `autonegotiate_variable` (String) Variable name
- `bandwidth_downstream` (Number) Interface downstream bandwidth capacity, in kbps
  - Range: `1`-`2147483647`
- `bandwidth_downstream_variable` (String) Variable name
- `bandwidth_upstream` (Number) Interface upstream bandwidth capacity, in kbps
  - Range: `1`-`2147483647`
- `bandwidth_upstream_variable` (String) Variable name
- `cellular_interface_name` (String) Cellular Interface Name <0-1>
- `cellular_interface_name_variable` (String) Variable name
- `clear_dont_fragment_bit` (Boolean) Clear don't fragment bit
  - Default value: `false`
- `clear_dont_fragment_bit_variable` (String) Variable name
- `core_region` (String) Enable core region
  - Choices: `core`, `core-shared`
  - Default value: `core`
- `core_region_variable` (String) Variable name
- `enable_core_region` (Boolean) Enable core region
  - Default value: `false`
- `enable_core_region_variable` (String) Variable name
- `interface_description` (String) Interface description
- `interface_description_variable` (String) Variable name
- `ip_directed_broadcast` (Boolean) IP Directed-Broadcast
  - Default value: `false`
- `ip_directed_broadcast_variable` (String) Variable name
- `ip_mtu` (Number) Interface MTU <68...1500>, in bytes
  - Range: `68`-`1500`
  - Default value: `1428`
- `ip_mtu_variable` (String) Variable name
- `ipv4_access_lists` (Attributes List) Apply ACL (see [below for nested schema](#nestedatt--ipv4_access_lists))
- `ipv4_dhcp_helper` (Set of String) List of DHCP server addresses
- `ipv4_dhcp_helper_variable` (String) Variable name
- `ipv6_access_lists` (Attributes List) Apply IPv6 access list (see [below for nested schema](#nestedatt--ipv6_access_lists))
- `nat` (Boolean) Network Address Translation on this interface
  - Default value: `false`
- `nat_block_icmp_error` (Boolean) Block inbound ICMP error messages
  - Default value: `true`
- `nat_block_icmp_error_variable` (String) Variable name
- `nat_port_forwards` (Attributes List) Set port-forwarding rules for NAT on this interface (see [below for nested schema](#nestedatt--nat_port_forwards))
- `nat_refresh_mode` (String) Set NAT refresh mode
  - Choices: `outbound`, `bi-directional`
  - Default value: `outbound`
- `nat_refresh_mode_variable` (String) Variable name
- `nat_response_to_ping` (Boolean) Respond to ping requests to NAT interface ip address from the public side
  - Default value: `false`
- `nat_response_to_ping_variable` (String) Variable name
- `nat_tcp_timeout` (Number) Set NAT TCP session timeout, in minutes
  - Range: `1`-`8947`
  - Default value: `60`
- `nat_tcp_timeout_variable` (String) Variable name
- `nat_udp_timeout` (Number) Set NAT UDP session timeout, in minutes
  - Range: `1`-`8947`
  - Default value: `1`
- `nat_udp_timeout_variable` (String) Variable name
- `per_tunnel_qos` (Boolean) Per-tunnel Qos
  - Default value: `false`
- `per_tunnel_qos_aggregator` (Boolean) Per-tunnel QoS Aggregator
  - Default value: `false`
- `per_tunnel_qos_aggregator_variable` (String) Variable name
- `per_tunnel_qos_variable` (String) Variable name
- `pmtu_discovery` (Boolean) Path MTU Discovery
  - Default value: `false`
- `pmtu_discovery_variable` (String) Variable name
- `policers` (Attributes List) Enable policer (see [below for nested schema](#nestedatt--policers))
- `qos_adaptive_bandwidth_downstream` (Number) Adaptive QoS default downstream bandwidth
  - Range: `8`-`100000000`
- `qos_adaptive_bandwidth_downstream_variable` (String) Variable name
- `qos_adaptive_bandwidth_upstream` (Number) Adaptive QoS default upstream bandwidth
  - Range: `8`-`100000000`
- `qos_adaptive_bandwidth_upstream_variable` (String) Variable name
- `qos_adaptive_max_downstream` (Number) Downstream max bandwidth limit
  - Range: `8`-`100000000`
- `qos_adaptive_max_downstream_variable` (String) Variable name
- `qos_adaptive_max_upstream` (Number) Upstream max bandwidth limit
  - Range: `8`-`100000000`
- `qos_adaptive_max_upstream_variable` (String) Variable name
- `qos_adaptive_min_downstream` (Number) Downstream min bandwidth limit
  - Range: `8`-`100000000`
- `qos_adaptive_min_downstream_variable` (String) Variable name
- `qos_adaptive_min_upstream` (Number) Upstream min bandwidth limit
  - Range: `8`-`100000000`
- `qos_adaptive_min_upstream_variable` (String) Variable name
- `qos_adaptive_period` (Number) Periodic timer for adaptive QoS in minutes
  - Range: `1`-`720`
  - Default value: `15`
- `qos_adaptive_period_variable` (String) Variable name
- `qos_map` (String) Name of QoS map
- `qos_map_variable` (String) Variable name
- `qos_map_vpn` (String) Name of VPN QoS map
- `qos_map_vpn_variable` (String) Variable name
- `secondary_region` (String) Enable secondary region
  - Choices: `off`, `secondary-only`, `secondary-shared`
  - Default value: `off`
- `secondary_region_variable` (String) Variable name
- `shaping_rate` (Number) 1ge  interfaces: [0..1000000]kbps; 10ge interfaces: [0..10000000]kbps
  - Range: `8`-`100000000`
- `shaping_rate_variable` (String) Variable name
- `shutdown` (Boolean) Administrative state
  - Default value: `true`
- `shutdown_variable` (String) Variable name
- `static_arps` (Attributes List) Configure static ARP entries (see [below for nested schema](#nestedatt--static_arps))
- `static_ingress_qos` (Number) Static ingress QoS for the port
  - Range: `0`-`7`
- `static_ingress_qos_variable` (String) Variable name
- `tcp_mss` (Number) TCP MSS on SYN packets, in bytes
  - Range: `552`-`1960`
- `tcp_mss_variable` (String) Variable name
- `tloc_extension` (String) Extends a local TLOC to a remote node only for vpn 0
- `tloc_extension_variable` (String) Variable name
- `tracker` (Set of String) Enable tracker for this interface
- `tracker_variable` (String) Variable name
- `tunnel_interface_allow_all` (Boolean) Allow all traffic. Overrides all other allow-service options if allow-service all is set
  - Default value: `false`
- `tunnel_interface_allow_all_variable` (String) Variable name
- `tunnel_interface_allow_bgp` (Boolean) Allow/deny BGP
  - Default value: `false`
- `tunnel_interface_allow_bgp_variable` (String) Variable name
- `tunnel_interface_allow_dhcp` (Boolean) Allow/Deny DHCP
  - Default value: `true`
- `tunnel_interface_allow_dhcp_variable` (String) Variable name
- `tunnel_interface_allow_dns` (Boolean) Allow/Deny DNS
  - Default value: `true`
- `tunnel_interface_allow_dns_variable` (String) Variable name
- `tunnel_interface_allow_https` (Boolean) Allow/Deny Https
  - Default value: `true`
- `tunnel_interface_allow_https_variable` (String) Variable name
- `tunnel_interface_allow_icmp` (Boolean) Allow/Deny ICMP
  - Default value: `true`
- `tunnel_interface_allow_icmp_variable` (String) Variable name
- `tunnel_interface_allow_netconf` (Boolean) Allow/Deny NETCONF
  - Default value: `false`
- `tunnel_interface_allow_netconf_variable` (String) Variable name
- `tunnel_interface_allow_ntp` (Boolean) Allow/Deny NTP
  - Default value: `false`
- `tunnel_interface_allow_ntp_variable` (String) Variable name
- `tunnel_interface_allow_ospf` (Boolean) Allow/Deny OSPF
  - Default value: `false`
- `tunnel_interface_allow_ospf_variable` (String) Variable name
- `tunnel_interface_allow_snmp` (Boolean) Allow/Deny SNMP
  - Default value: `false`
- `tunnel_interface_allow_snmp_variable` (String) Variable name
- `tunnel_interface_allow_ssh` (Boolean) Allow/Deny SSH
  - Default value: `false`
- `tunnel_interface_allow_ssh_variable` (String) Variable name
- `tunnel_interface_allow_stun` (Boolean) Allow/Deny STUN
  - Default value: `false`
- `tunnel_interface_allow_stun_variable` (String) Variable name
- `tunnel_interface_bind_loopback_tunnel` (String) Bind loopback tunnel interface to a physical interface
- `tunnel_interface_bind_loopback_tunnel_variable` (String) Variable name
- `tunnel_interface_border` (Boolean) Set TLOC as border TLOC
  - Default value: `false`
- `tunnel_interface_border_variable` (String) Variable name
- `tunnel_interface_carrier` (String) Set carrier for TLOC
  - Choices: `default`, `carrier1`, `carrier2`, `carrier3`, `carrier4`, `carrier5`, `carrier6`, `carrier7`, `carrier8`
  - Default value: `default`
- `tunnel_interface_carrier_variable` (String) Variable name
- `tunnel_interface_clear_dont_fragment` (Boolean) Enable clear dont fragment (Currently Only SDWAN Tunnel Interface)
  - Default value: `false`
- `tunnel_interface_clear_dont_fragment_variable` (String) Variable name
- `tunnel_interface_color` (String) Set color for TLOC
  - Choices: `default`, `mpls`, `metro-ethernet`, `biz-internet`, `public-internet`, `lte`, `3g`, `red`, `green`, `blue`, `gold`, `silver`, `bronze`, `custom1`, `custom2`, `custom3`, `private1`, `private2`, `private3`, `private4`, `private5`, `private6`
  - Default value: `default`
- `tunnel_interface_color_restrict` (Boolean) Restrict this TLOC behavior
  - Default value: `false`
- `tunnel_interface_color_restrict_variable` (String) Variable name
- `tunnel_interface_color_variable` (String) Variable name
- `tunnel_interface_control_connections` (Boolean) Allow Control Connection
  - Default value: `true`
- `tunnel_interface_control_connections_variable` (String) Variable name
- `tunnel_interface_encapsulations` (Attributes List) Encapsulation for TLOC (see [below for nested schema](#nestedatt--tunnel_interface_encapsulations))
- `tunnel_interface_exclude_controller_group_list` (Set of Number) Exclude the following controller groups defined in this list
- `tunnel_interface_exclude_controller_group_list_variable` (String) Variable name
- `tunnel_interface_groups` (Set of Number) List of groups
- `tunnel_interface_groups_variable` (String) Variable name
- `tunnel_interface_hello_interval` (Number) Set time period of control hello packets <100..600000> milli seconds
  - Range: `100`-`600000`
  - Default value: `1000`
- `tunnel_interface_hello_interval_variable` (String) Variable name
- `tunnel_interface_hello_tolerance` (Number) Set tolerance of control hello packets <12..6000> seconds
  - Range: `12`-`6000`
  - Default value: `12`
- `tunnel_interface_hello_tolerance_variable` (String) Variable name
- `tunnel_interface_last_resort_circuit` (Boolean) Set TLOC as last resort
  - Default value: `false`
- `tunnel_interface_last_resort_circuit_variable` (String) Variable name
- `tunnel_interface_low_bandwidth_link` (Boolean) Set the interface as a low-bandwidth circuit
  - Default value: `false`
- `tunnel_interface_low_bandwidth_link_variable` (String) Variable name
- `tunnel_interface_max_control_connections` (Number) Set the maximum number of control connections for this TLOC
  - Range: `0`-`8`
- `tunnel_interface_max_control_connections_variable` (String) Variable name
- `tunnel_interface_nat_refresh_interval` (Number) Set time period of nat refresh packets <1...60> seconds
  - Range: `1`-`60`
  - Default value: `5`
- `tunnel_interface_nat_refresh_interval_variable` (String) Variable name
- `tunnel_interface_network_broadcast` (Boolean) Accept and respond to network-prefix-directed broadcasts)
  - Default value: `false`
- `tunnel_interface_network_broadcast_variable` (String) Variable name
- `tunnel_interface_port_hop` (Boolean) Disallow port hopping on the tunnel interface
  - Default value: `true`
- `tunnel_interface_port_hop_variable` (String) Variable name
- `tunnel_interface_tunnel_tcp_mss` (Number) Tunnel TCP MSS on SYN packets, in bytes
  - Range: `500`-`1460`
- `tunnel_interface_tunnel_tcp_mss_variable` (String) Variable name
- `tunnel_interface_vbond_as_stun_server` (Boolean) Put this wan interface in STUN mode only
  - Default value: `false`
- `tunnel_interface_vbond_as_stun_server_variable` (String) Variable name
- `tunnel_interface_vmanage_connection_preference` (Number) Set interface preference for control connection to vManage <0..8>
  - Range: `0`-`8`
  - Default value: `5`
- `tunnel_interface_vmanage_connection_preference_variable` (String) Variable name
- `tunnel_qos_mode` (String) Set tunnel QoS mode
  - Choices: `spoke`
- `tunnel_qos_mode_variable` (String) Variable name
- `write_rule` (String) Name of rewrite rule
- `write_rule_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the feature template
- `template_type` (String) The template type
- `version` (Number) The version of the feature template

<a id="nestedatt--ipv4_access_lists"></a>
### Nested Schema for `ipv4_access_lists`

Optional:

- `acl_name` (String) Name of access list
- `acl_name_variable` (String) Variable name
- `direction` (String) Direction
  - Choices: `in`, `out`
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--ipv6_access_lists"></a>
### Nested Schema for `ipv6_access_lists`

Optional:

- `acl_name` (String) Name of access list
- `acl_name_variable` (String) Variable name
- `direction` (String) Direction
  - Choices: `in`, `out`
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--nat_port_forwards"></a>
### Nested Schema for `nat_port_forwards`

Optional:

- `optional` (Boolean) Indicates if list item is considered optional.
- `port_end_range` (Number) Ending port of port range
  - Range: `0`-`65535`
- `port_start_range` (Number) Starting port of port range
  - Range: `0`-`65535`
- `private_ip_address` (String) Private IP Address to translate to
- `private_ip_address_variable` (String) Variable name
- `private_vpn` (Number) VPN in which private IP Address resides
  - Range: `0`-`65535`
- `private_vpn_variable` (String) Variable name
- `protocol` (String) Layer 4 protocol to apply port forwarding to
  - Choices: `tcp`, `udp`


<a id="nestedatt--policers"></a>
### Nested Schema for `policers`

Optional:

- `direction` (String) Direction
  - Choices: `in`, `out`
- `optional` (Boolean) Indicates if list item is considered optional.
- `policer_name` (String) Name of policer


<a id="nestedatt--static_arps"></a>
### Nested Schema for `static_arps`

Optional:

- `ip_address` (String) IP Address
- `ip_address_variable` (String) Variable name
- `mac` (String) MAC address
- `mac_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--tunnel_interface_encapsulations"></a>
### Nested Schema for `tunnel_interface_encapsulations`

Optional:

- `encapsulation` (String) Encapsulation
  - Choices: `gre`, `ipsec`
- `optional` (Boolean) Indicates if list item is considered optional.
- `preference` (Number) Set preference for TLOC
  - Range: `0`-`4294967295`
- `preference_variable` (String) Variable name
- `weight` (Number) Set weight for TLOC
  - Range: `1`-`255`
  - Default value: `1`
- `weight_variable` (String) Variable name

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_vpn_interface_cellular_feature_template.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```