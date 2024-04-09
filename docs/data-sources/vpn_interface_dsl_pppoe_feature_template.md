---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_vpn_interface_dsl_pppoe_feature_template Data Source - terraform-provider-sdwan"
subcategory: "Feature Templates"
description: |-
  This data source can read the VPN Interface DSL PPPoE feature template.
---

# sdwan_vpn_interface_dsl_pppoe_feature_template (Data Source)

This data source can read the VPN Interface DSL PPPoE feature template.

## Example Usage

```terraform
data "sdwan_vpn_interface_dsl_pppoe_feature_template" "example" {
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
- `adapt_period` (Number) Periodic timer for adaptive QoS in minutes
- `adapt_period_variable` (String) Variable name
- `adaptive_qos` (Boolean) Adaptive QoS
- `all` (Boolean) Allow all traffic. Overrides all other allow-service options if allow-service all is set
- `all_variable` (String) Variable name
- `authentication_type` (String) Authenticate remote on incoming call only
- `authentication_type_variable` (String) Variable name
- `bandwidth_downstream` (Number) Interface downstream bandwidth capacity, in kbps
- `bandwidth_downstream_variable` (String) Variable name
- `bandwidth_upstream` (Number) Interface upstream bandwidth capacity, in kbps
- `bandwidth_upstream_variable` (String) Variable name
- `bgp` (Boolean) Allow/deny BGP
- `bgp_variable` (String) Variable name
- `bind_loopback_tunnel` (String) Bind loopback tunnel interface to a physical interface
- `bind_loopback_tunnel_variable` (String) Variable name
- `block_icmp_error` (Boolean) Block inbound ICMP error messages
- `block_icmp_error_variable` (String) Variable name
- `border` (Boolean) Set TLOC as border TLOC
- `border_variable` (String) Variable name
- `carrier` (String) Set carrier for TLOC
- `carrier_variable` (String) Variable name
- `chap_hostname` (String) CHAP Hostname
- `chap_hostname_variable` (String) Variable name
- `chap_password` (String) Specify ppp chap authentication Password
- `chap_password_variable` (String) Variable name
- `clear_dont_fragment` (Boolean) Enable clear dont fragment (Currently Only SDWAN Tunnel Interface)
- `clear_dont_fragment_variable` (String) Variable name
- `color` (String) Set color for TLOC
- `color_variable` (String) Variable name
- `control_connections` (Boolean) Allow Control Connection
- `control_connections_variable` (String) Variable name
- `core_region` (String) Enable core region
- `core_region_variable` (String) Variable name
- `description` (String) The description of the feature template
- `device_types` (Set of String) List of supported device types
- `dhcp` (Boolean) Allow/Deny DHCP
- `dhcp_variable` (String) Variable name
- `dialer_address_negotiated` (Boolean) Dialer IP Negotiated
- `dialer_address_negotiated_variable` (String) Variable name
- `dialer_pool_number` (Number) Dialer pool number
- `dialer_pool_number_variable` (String) Variable name
- `dns` (Boolean) Allow/Deny DNS
- `dns_variable` (String) Variable name
- `enable_core_region` (Boolean) Enable core region
- `enable_core_region_variable` (String) Variable name
- `encap` (Number) Encapsulation VLAN id
- `encap_variable` (String) Variable name
- `encapsulation` (Attributes List) Encapsulation for TLOC (see [below for nested schema](#nestedatt--encapsulation))
- `ethernet_interface_name` (String) Ethernet Interface/Sub Interface Name including sub interface number
- `ethernet_interface_name_variable` (String) Variable name
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
- `interface_description` (String) Interface description
- `interface_description_variable` (String) Variable name
- `ip_directed_broadcast` (Boolean) IP Directed-Broadcast
- `ip_directed_broadcast_variable` (String) Variable name
- `ip_mtu` (Number) Interface MTU <576..2000>, in bytes
- `ip_mtu_variable` (String) Variable name
- `last_resort_circuit` (Boolean) Set TLOC as last resort
- `last_resort_circuit_variable` (String) Variable name
- `low_bandwidth_link` (Boolean) Set the interface as a low-bandwidth circuit
- `low_bandwidth_link_variable` (String) Variable name
- `max_control_connections` (Number) Set the maximum number of control connections for this TLOC
- `max_control_connections_variable` (String) Variable name
- `nat` (Boolean) Network Address Translation on this interface
- `nat_refresh_interval` (Number) Set time period of nat refresh packets <1...60> seconds
- `nat_refresh_interval_variable` (String) Variable name
- `nat_variable` (String) Variable name
- `netconf` (Boolean) Allow/Deny NETCONF
- `netconf_variable` (String) Variable name
- `network_broadcast_1` (Boolean) Accept and respond to network-prefix-directed broadcasts)
- `network_broadcast_1_variable` (String) Variable name
- `network_broadcast_2` (Boolean) Accept and respond to network-prefix-directed broadcasts)
- `network_broadcast_2_variable` (String) Variable name
- `ospf` (Boolean) Allow/Deny OSPF
- `ospf_variable` (String) Variable name
- `pap_outbound_password` (String) Specify ppp pap authentication Password
- `pap_outbound_password_variable` (String) Variable name
- `pap_password` (Boolean) PAP outbound Password
- `pap_username` (String) PAP outbound Sent Username
- `pap_username_variable` (String) Variable name
- `per_tunnel_qos` (Boolean) Per-tunnel Qos
- `per_tunnel_qos_aggregator` (Boolean) Per-tunnel QoS Aggregator
- `per_tunnel_qos_aggregator_variable` (String) Variable name
- `per_tunnel_qos_variable` (String) Variable name
- `policer` (Attributes List) Enable policer (see [below for nested schema](#nestedatt--policer))
- `port_forward` (Attributes List) Set port-forwarding rules for NAT on this interface (see [below for nested schema](#nestedatt--port_forward))
- `port_hop` (Boolean) Disallow port hopping on the tunnel interface
- `port_hop_variable` (String) Variable name
- `ppp_authentication_protocol` (String) PPP Link Authentication Protocol
- `ppp_authentication_protocol_pap` (Boolean) PPP Authentication Protocol PAP
- `ppp_authentication_protocol_pap_variable` (String) Variable name
- `ppp_authentication_protocol_variable` (String) Variable name
- `ppp_maximum_payload` (Number) Maximum MRU to be negotiated during PPP LCP negotiation
- `ppp_maximum_payload_variable` (String) Variable name
- `qos_map` (String) Name of QoS map
- `qos_map_variable` (String) Variable name
- `refresh_mode` (String) Set NAT refresh mode
- `refresh_mode_variable` (String) Variable name
- `respond_to_ping` (Boolean) Respond to ping requests to NAT interface ip address from the public side
- `respond_to_ping_variable` (String) Variable name
- `restrict` (Boolean) Restrict this TLOC behavior
- `restrict_variable` (String) Variable name
- `secondary_region` (String) Enable secondary region
- `secondary_region_variable` (String) Variable name
- `shaping_rate` (Number) 1ge  interfaces: [0..1000000]kbps; 10ge interfaces: [0..10000000]kbps
- `shaping_rate_downstream_default` (Number) Adaptive QoS default downstream bandwidth
- `shaping_rate_downstream_default_variable` (String) Variable name
- `shaping_rate_downstream_max` (Number) Downstream max bandwidth limit
- `shaping_rate_downstream_max_variable` (String) Variable name
- `shaping_rate_downstream_min` (Number) Downstream min bandwidth limit
- `shaping_rate_downstream_min_variable` (String) Variable name
- `shaping_rate_upstream_default` (Number) Adaptive QoS default upstream bandwidth
- `shaping_rate_upstream_default_variable` (String) Variable name
- `shaping_rate_upstream_max` (Number) Upstream max bandwidth limit
- `shaping_rate_upstream_max_variable` (String) Variable name
- `shaping_rate_upstream_min` (Number) Upstream min bandwidth limit
- `shaping_rate_upstream_min_variable` (String) Variable name
- `shaping_rate_variable` (String) Variable name
- `shutdown` (Boolean) Administrative state
- `shutdown_variable` (String) Variable name
- `snmp` (Boolean) Allow/Deny SNMP
- `snmp_variable` (String) Variable name
- `ssh` (Boolean) Allow/Deny SSH
- `ssh_variable` (String) Variable name
- `stun` (Boolean) Allow/Deny STUN
- `stun_variable` (String) Variable name
- `tcp_mss` (Number) TCP MSS on SYN packets, in bytes
- `tcp_mss_variable` (String) Variable name
- `tcp_timeout` (Number) Set NAT TCP session timeout, in minutes
- `tcp_timeout_variable` (String) Variable name
- `template_type` (String) The template type
- `tloc_extension` (String) Extends a local TLOC to a remote node only for vpn 0
- `tloc_extension_variable` (String) Variable name
- `tracker` (Set of String) Enable tracker for this interface
- `tracker_variable` (String) Variable name
- `tunnel_tcp_mss` (Number) Tunnel TCP MSS on SYN packets, in bytes
- `tunnel_tcp_mss_variable` (String) Variable name
- `udp_timeout` (Number) Set NAT UDP session timeout, in minutes
- `udp_timeout_variable` (String) Variable name
- `unnumbered_loopback_interface` (String) Dialer IP Unnumbered Loopback interface name
- `unnumbered_loopback_interface_variable` (String) Variable name
- `vbond_as_stun_server` (Boolean) Put this wan interface in STUN mode only
- `vbond_as_stun_server_variable` (String) Variable name
- `vdsl` (Attributes List) vdsl (see [below for nested schema](#nestedatt--vdsl))
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


<a id="nestedatt--encapsulation"></a>
### Nested Schema for `encapsulation`

Read-Only:

- `encapsulation_type` (String) Encapsulation
- `optional` (Boolean) Indicates if list item is considered optional.
- `preference` (Number) Set preference for TLOC
- `preference_variable` (String) Variable name
- `weight` (Number) Set weight for TLOC
- `weight_variable` (String) Variable name


<a id="nestedatt--policer"></a>
### Nested Schema for `policer`

Read-Only:

- `direction` (String) Direction
- `optional` (Boolean) Indicates if list item is considered optional.
- `policer_name` (String) Name of policer


<a id="nestedatt--port_forward"></a>
### Nested Schema for `port_forward`

Read-Only:

- `optional` (Boolean) Indicates if list item is considered optional.
- `port_end_range` (Number) Ending port of port range
- `port_start_range` (Number) Starting port of port range
- `private_ip_address` (String) Private IP Address to translate to
- `private_ip_address_variable` (String) Variable name
- `private_vpn` (Number) VPN in which private IP Address resides
- `private_vpn_variable` (String) Variable name
- `protocol` (String) Layer 4 protocol to apply port forwarding to


<a id="nestedatt--vdsl"></a>
### Nested Schema for `vdsl`

Read-Only:

- `controller_vdsl_slot` (String) Set module slot/subslot/port number
- `controller_vdsl_slot_variable` (String) Variable name
- `mode_adsl1` (Boolean) Set VDSL operating mode to adsl1
- `mode_adsl1_variable` (String) Variable name
- `mode_adsl2` (Boolean) Set VDSL operating mode to adsl2
- `mode_adsl2_variable` (String) Variable name
- `mode_adsl2plus` (Boolean) Set VDSL operating mode to adsl2plus
- `mode_adsl2plus_variable` (String) Variable name
- `mode_ansi` (Boolean) Set VDSL operating mode to ansi
- `mode_ansi_variable` (String) Variable name
- `mode_vdsl2` (Boolean) Set VDSL operating mode to vdsl2
- `mode_vdsl2_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `sra` (Boolean) Seamless rate adaption
- `sra_variable` (String) Variable name
- `vdsl_modem_configuration` (String) Set module slot/subslot/port number
- `vdsl_modem_configuration_variable` (String) Variable name