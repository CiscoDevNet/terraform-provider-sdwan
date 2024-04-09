---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_vpn_interface_t1_e1_serial_feature_template Data Source - terraform-provider-sdwan"
subcategory: "Feature Templates"
description: |-
  This data source can read the VPN Interface T1 E1 Serial feature template.
---

# sdwan_vpn_interface_t1_e1_serial_feature_template (Data Source)

This data source can read the VPN Interface T1 E1 Serial feature template.

## Example Usage

```terraform
data "sdwan_vpn_interface_t1_e1_serial_feature_template" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The id of the feature template
- `name` (String) The name of the feature template

### Read-Only

- `access_list` (Attributes List) Apply ACL (see [below for nested schema](#nestedatt--access_list))
- `administrative_shutdown` (Boolean) Administrative state
- `administrative_shutdown_variable` (String) Variable name
- `all` (Boolean) Allow all traffic. Overrides all other allow-service options if allow-service all is set
- `all_variable` (String) Variable name
- `bgp` (Boolean) Allow/deny BGP
- `bgp_variable` (String) Variable name
- `bind_loopback_tunnel` (String) Bind loopback tunnel interface to a physical interface
- `bind_loopback_tunnel_variable` (String) Variable name
- `border` (Boolean) Set TLOC as border TLOC
- `border_variable` (String) Variable name
- `carrier` (String) Set carrier for TLOC
- `carrier_variable` (String) Variable name
- `clear_dont_fragment_bit` (Boolean) Clear don't fragment bit
- `clear_dont_fragment_bit_variable` (String) Variable name
- `clock_rate` (String) Set preference for interface Clock speed
- `clock_rate_variable` (String) Variable name
- `color` (String) Set color for TLOC
- `color_variable` (String) Variable name
- `core_region` (String) Enable core region
- `core_region_variable` (String) Variable name
- `description` (String) The description of the feature template
- `device_types` (Set of String) List of supported device types
- `dhcp` (Boolean) Allow/Deny DHCP
- `dhcp_variable` (String) Variable name
- `dns` (Boolean) Allow/Deny DNS
- `dns_variable` (String) Variable name
- `enable_clear_don` (Boolean) Enable clear dont fragment (Currently Only SDWAN Tunnel Interface)
- `enable_clear_don_variable` (String) Variable name
- `enable_core_region` (Boolean) Enable core region
- `enable_core_region_variable` (String) Variable name
- `encapsulation` (String) Configure Encapsulation for interface
- `encapsulation_for_tloc` (Attributes List) Encapsulation for TLOC (see [below for nested schema](#nestedatt--encapsulation_for_tloc))
- `encapsulation_variable` (String) Variable name
- `exclude_controller_group_list` (Set of Number) Exclude the following controller groups defined in this list
- `exclude_controller_group_list_variable` (String) Variable name
- `groups` (Set of Number) List of groups
- `groups_variable` (String) Variable name
- `hello_interval` (Number) Set time period of control hello packets <100..600000> milli seconds
- `hello_interval_variable` (String) Variable name
- `hello_tolerance` (Number) Set tolerance of control hello packets <12..6000> seconds
- `hello_tolerance_variable` (String) Variable name
- `https` (Boolean) Allow/Deny Https
- `https_variable` (String) Variable name
- `icmp` (Boolean) Allow/Deny ICMP
- `icmp_variable` (String) Variable name
- `interface_bandwidth_capacity` (Number) Interface bandwidth capacity, in kbps
- `interface_bandwidth_capacity_variable` (String) Variable name
- `interface_description` (String) Interface description
- `interface_description_variable` (String) Variable name
- `interface_downstream_bandwidth_capacity` (Number) Interface downstream bandwidth capacity, in kbps
- `interface_downstream_bandwidth_capacity_variable` (String) Variable name
- `ip_mtu` (Number) Interface MTU <68...2000>, in bytes
- `ip_mtu_variable` (String) Variable name
- `ipv4_address` (String) Assign IPv4 address
- `ipv4_address_variable` (String) Variable name
- `ipv6_access_lists` (Attributes List) Apply IPv6 access list (see [below for nested schema](#nestedatt--ipv6_access_lists))
- `ipv6_address` (String) Assign IPv6 address
- `ipv6_address_variable` (String) Variable name
- `last_resort_circuit` (Boolean) Set TLOC as last resort
- `last_resort_circuit_variable` (String) Variable name
- `link_autonegotiate` (Boolean) Link autonegotiation
- `link_autonegotiate_variable` (String) Variable name
- `low_bandwidth_link` (Boolean) Set the interface as a low-bandwidth circuit
- `low_bandwidth_link_variable` (String) Variable name
- `max_control_connections` (Number) Set the maximum number of control connections for this TLOC
- `max_control_connections_variable` (String) Variable name
- `nat_refresh_interval` (Number) Set time period of nat refresh packets <1...60> seconds
- `nat_refresh_interval_variable` (String) Variable name
- `netconf` (Boolean) Allow/Deny NETCONF
- `netconf_variable` (String) Variable name
- `network_broadcast_1` (Boolean) Accept and respond to network-prefix-directed broadcasts)
- `network_broadcast_1_variable` (String) Variable name
- `network_broadcast_2` (Boolean) Accept and respond to network-prefix-directed broadcasts)
- `network_broadcast_2_variable` (String) Variable name
- `ntp` (Boolean) Allow/Deny NTP
- `ntp_variable` (String) Variable name
- `ospf` (Boolean) Allow/Deny OSPF
- `ospf_variable` (String) Variable name
- `per_tunnel_qos` (Boolean) Per-tunnel Qos
- `per_tunnel_qos_aggregator` (Boolean) Per-tunnel QoS Aggregator
- `per_tunnel_qos_aggregator_variable` (String) Variable name
- `per_tunnel_qos_variable` (String) Variable name
- `pmtu_discovery` (Boolean) Path MTU Discovery
- `pmtu_discovery_variable` (String) Variable name
- `port_hop` (Boolean) Disallow port hopping on the tunnel interface
- `port_hop_variable` (String) Variable name
- `qos_map` (String) Name of QoS map
- `qos_map_variable` (String) Variable name
- `restrict` (Boolean) Restrict this TLOC behavior
- `restrict_variable` (String) Variable name
- `secondary_region` (String) Enable secondary region
- `secondary_region_variable` (String) Variable name
- `serial_interface_name` (String) Serial Interface Name - slot/subslot/port:channel-group for T1/E1, slot/subslot/port for NIM-1T
- `serial_interface_name_variable` (String) Variable name
- `shaping_rate` (Number) 1ge  interfaces: [0..1000000]kbps; 10ge interfaces: [0..10000000]kbps
- `shaping_rate_variable` (String) Variable name
- `snmp` (Boolean) Allow/Deny SNMP
- `snmp_variable` (String) Variable name
- `ssh` (Boolean) Allow/Deny SSH
- `ssh_variable` (String) Variable name
- `static_ingress_qos` (Number) Static ingress QoS for the port
- `static_ingress_qos_variable` (String) Variable name
- `stun` (Boolean) Allow/Deny STUN
- `stun_variable` (String) Variable name
- `tcp_mss` (Number) TCP MSS on SYN packets, in bytes
- `tcp_mss_variable` (String) Variable name
- `template_type` (String) The template type
- `tloc_extension` (String) Extends a local TLOC to a remote node only for vpn 0
- `tloc_extension_variable` (String) Variable name
- `tunnel_tcp_mss` (Number) Tunnel TCP MSS on SYN packets, in bytes
- `tunnel_tcp_mss_variable` (String) Variable name
- `vbond_as_stun_server` (Boolean) Put this wan interface in STUN mode only
- `vbond_as_stun_server_variable` (String) Variable name
- `version` (Number) The version of the feature template
- `vmanage_connection_preference` (Number) Set interface preference for control connection to vManage <0..8>
- `vmanage_connection_preference_variable` (String) Variable name
- `vpn_qos_map` (String) Name of VPN QoS map
- `vpn_qos_map_variable` (String) Variable name
- `write_rule` (String) Name of rewrite rule
- `write_rule_variable` (String) Variable name

<a id="nestedatt--access_list"></a>
### Nested Schema for `access_list`

Read-Only:

- `acl_name` (String) Name of access list
- `acl_name_variable` (String) Variable name
- `direction` (String) Direction
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--encapsulation_for_tloc"></a>
### Nested Schema for `encapsulation_for_tloc`

Read-Only:

- `encapsulation_type` (String) Encapsulation
- `optional` (Boolean) Indicates if list item is considered optional.
- `preference` (Number) Set preference for TLOC
- `preference_variable` (String) Variable name
- `weight` (Number) Set weight for TLOC
- `weight_variable` (String) Variable name


<a id="nestedatt--ipv6_access_lists"></a>
### Nested Schema for `ipv6_access_lists`

Read-Only:

- `acl_name` (String) Name of access list
- `acl_name_variable` (String) Variable name
- `direction` (String) Direction
- `optional` (Boolean) Indicates if list item is considered optional.