---
layout: "sdwan"
page_title: "SD-WAN: sdwan_vpn_interface_feature_template"
sidebar_current: "docs-sdwan-data-source-vpn_interface_feature_template"
description: |-
  Data source for SD-WAN VPN interface Feature Template
---

# sdwan_vpn_interface_feature_template Data Source #
Data source for SD-WAN VPN interface Feature Templates

## Example Usage ##

```hcl

data "sdwan_vpn_interface_feature_template" "name" {
    template_name = "Sample_Vedge_VPN_Interface"
}

```

## Argument Reference ##

* `template_name` - (Required) Unique VPN interface Type Feature Template Name


## Attribute Reference ##

* `template_description` - Long Description of VPN interface type Feature Template
* `device_type` - Type of device which supports VPN interface Feature Template
* `template_type` - Template type for the VPN Interface type Feature Template
* `template_min_version` - Minimum Version for the VPN interface Feature template
* `template_definition` - Configuration for VPN interface type Feature Template, (see [below for nested schema](#nestedblock--template_definition))
* `factory_default` - Boolean flag indicating whether VPN interface type Feature Template is factory default or not
* `attached_masters_count` - Number of Device Templates attached to VPN interface type Feature Template
* `devices_attached` - Number of Devices attached to VPN interface type Feature Template
* `last_updated_by` - User who updated the VPN interface type Feature Template latest
* `last_updated_on` - Time for the latest Update of VPN interface type Feature Template
* `g_template_class` - Template Class type for VPN interface type Feature Template
* `template_id` - Unique ID for VPN interface type Feature Template
* `config_type` - Type of configuration for VPN interface type Feature Template
* `created_on` - Time of creation of VPN interface type Feature Template
* `rid` - rid for the VPN interface type Feature Template
* `feature` - Feature for the VPN interface type Feature Template

<a id="nestedblock--template_definition"></a>

## Nested Schema for template_definition
* `vpn_interface_basic` - Basic configuration for the  VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_basic))
* `vpn_interface_tunnel` - Tunnel configuration for the  VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_tunnel))
* `vpn_interface_nat` - Tunnel configuration for the  VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_nat))
* `vpn_interface_vrrp` - VRRP(Virtual Router Redundancy Protocol) configuration for the  VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_vrrp))
* `vpn_interface_acl_qos` - ACL(Apply Access Lists) and QoS  configuration for the  VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_acl_qos))
* `vpn_interface_arp` - ARP(Address Resolution Protocol) configuration for the  VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_arp))
* `vpn_interface_8021x` - IEEE 802.1X Authentication configuration for the  Vedge VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_8021x))
* `vpn_interface_trustsec` - TrustSec configuration for the Cisco VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_trustsec)) 
* `vpn_interface_advanced` - Advanced configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_advanced))


<a id="nestedblock--vpn_interface_basic"></a>

## Nested Schema for vpn_interface_basic
* `shutdown` - Shutdown flag for VPN interface type Feature Template
* `description` - Interface Description for VPN interface type Feature Template
* `ipv4` - IPv4 configuration for VPN interface type Feature Template, (see [below for nested schema](#nestedblock--ipv4))
* `ipv6` - IPv6 configuration for VPN interface type Feature Template, (see [below for nested schema](#nestedblock--ipv6))
* `block_non_source_ip` - Flag indicating whether forwarded traffic matches with interface's IP prefix range for VPN interface type Feature Template
* `bandwidth_upstream` - Value of bandwidth above which to generate notifications in case of transmitted traffic for VPN interface type Feature Template
* `bandwidth_downstream` - Value of bandwidth above which to generate notifications in case of received traffic for VPN interface type Feature Template


<a id="nestedblock--vpn_interface_tunnel"></a>

## Nested Schema for vpn_interface_tunnel
* `per_tunnel_qos` - Flag indicating whether per_tunnel_qos is enabled or not for VPN interface type Feature Template
* `per_tunnel_qos_aggregator` - Flag indicating whether per_tunnel_qos_aggregator is enabled or not for VPN interface type Feature Template
* `tunnels_bandwidth_percent` - Value of tunnel bandwith percentage for VPN interface type Feature Template
* `color` - TLOC color value for VPN interface type Feature Template
* `restrict` - Restrict flag for VPN interface type Feature Template
* `groups` - Value of Groups in tunnle for VPN interface type Feature Template
* `border` - Flag indicating whether Border is enabled or not for VPN interface type Feature Template
* `control_connection` - Flag indicating TLOC connection is enabled or not for VPN interface type Feature Template 
* `maximum_control_connections` - Maximum number of vSamart controllers that WAN tunnel interface can connect for VPN interface type Feature Template 
* `vbond_as_stun_server` - Flag indicating whether Session Traversal Utilities for NAT(STUN) is enabled or not for VPN interface type Feature Template
* `exclude_controller_group_list` - List of vSmart controllers that the tunnel interface is not allowed to connect for VPN interface type Feature Template
* `vmanage_connection_preference` - Flag indicating preference for using a tunnel interface to exchange control traffic with the vManage NMS for VPN interface type Feature Template
* `port_hop` - Flag indicating whether Port Hop is enabled or not for VPN interface type Feature Template
* `low_bandwidth_link` - Flag indicating whether tunnle is set as a low-bandwidth link for VPN interface type Feature Template
* `allow_service` - Nested block for allow or disallow various services for VPN interface type Feature Template(see [below for nested schema](#nestedblock--allow_service))
* `encapsulation` - Encapsulation configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--encapsulation))


<a id="nestedblock--vpn_interface_nat"></a>

## Nested Schema for vpn_interface_nat
* `ipv4` - NAT IPv4 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--natv4))
* `ipv6` - NAT IPv6 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--natv6))



<a id="nestedblock--vpn_interface_vrrp"></a>

## Nested Schema for vpn_interface_vrrp
* `ipv4` - VRRP IPv4 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--vrrpv4))
* `ipv6` - VRRP IPv6 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--vrrpv6))


<a id="nestedblock--vpn_interface_acl_qos"></a>

## Nested Schema for vpn_interface_acl_qos
* `adapt_period` - Value of Adapt Period for Cisco VPN interface type Feature Template
* `shaping_rate_upstream_min` - Minimum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template
* `shaping_rate_upstream_max` - Maximum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template
* `shaping_rate_upstream_default` - Default value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template
* `shaping_rate_downstream_min` - Minimum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template
* `shaping_rate_downstream_max` - Maximum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template
* `shaping_rate_downstream_default` - Default value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template
* `shaping_rate` - Value of aggregate traffic transmission rate on the VPN interface for VPN interface type Feature Template
* `qos_map` - Name of the QoS map to apply to packets being transmitted out the interface for VPN interface type Feature Template
* `rewrite_rule` - Name of the rewrite rule to apply on the interface for VPN interface type Feature Template
* `ipv4_ingress_access_list` - Name of the access list to apply to IPv4 packets being received on the interface for VPN interface type Feature Template
* `ipv4_egress_access_list` - Name of the access list to apply to IPv4 packets being transmitted on the interface for VPN interface type Feature Template
* `ipv6_ingress_access_list` - Name of the access list to apply to IPv6 packets being received on the interface for VPN interface type Feature Template
* `ipv6_egress_access_list` - Name of the access list to apply to IPv6 packets being transmitted on the interface for VPN interface type Feature Template
* `ingress_policer_name` - Name of the policer to apply to packets received on the interface for VPN interface type Feature Template
* `egress_policer_name` - Name of the policer to apply to packets being transmitted on the interface for Vedge VPN interface type Feature Template


<a id="nestedblock--vpn_interface_arp"></a>

## Nested Schema for vpn_interface_arp
* `ip_address` - IP address for the ARP entry in dotted decimal notation or as a fully qualified host name for VPN interface type Feature Template
* `mac_address` - MAC address in colon-separated hexadecimal notation for VPN interface type Feature Template



<a id="nestedblock--vpn_interface_8021x"></a>

## Nested Schema for vpn_interface_8021x
* `radius_server` - The tag of the RADIUS server to use for 802.1X authentication for Vedge VPN interface type Feature Template
* `account_interim_interval` - Value of how often to send 802.1X interim accounting updates to the RADIUS server for Vedge VPN interface type Feature Template
* `nas_identifier` - The NAS identifier of the local router for Vedge VPN interface type Feature Template
* `nas_ip` - The NAS IP address of the local router for Vedge VPN interface type Feature Template
* `wake_on_lan` - The flag that enable client to be powered up when the router receives an Ethernet magic packet frame for Vedge VPN interface type Feature Template
* `control_direction` - Direction type of packets for Vedge VPN interface type Feature Template
* `reauthentication` - Value indicating how often to reauthenticate 802.1X clients for Vedge VPN interface type Feature Template
* `inactivity` - Value indicating how long to wait before revoking an 802.1X client's network access for Vedge VPN interface type Feature Template
* `host_mode` - Value indicating whether an 802.1X interface grants access to a single client or to multiple clients for Vedge VPN interface type Feature Template
* `advanced_options` - Advance configuration of vpn_interface_8021x for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--advanced_options))


<a id="nestedblock--vpn_interface_trustsec"></a>

## Nested Schema for vpn_interface_trustsec
* `enable_sgt_propagation` - Flag indicating whether SGT propagation is enabled or not for Cisco VPN interface type Feature Template
* `propagate` - Flag indicating whether propagate is enabled or not for Cisco VPN interface type Feature Template
* `security_group_tag` - Value of security group tag for Cisco VPN interface type Feature Template
* `trusted` - Flag indicating whether trusted is enable or not for Cisco VPN interface type Feature Template



<a id="nestedblock--vpn_interface_advanced"></a>

## Nested Schema for vpn_interface_advanced
* `duplex` - Value of duplex to run interface in auto, full, or half mode for VPN interface type Feature Template
* `mac_address` - Value of MAC address to associate with the interface, in colon-separated hexadecimal notation for VPN interface type Feature Template
* `ip_mtu` - Value of MAC address to associate with the interface, in colon-separated hexadecimal notation for VPN interface type Feature Template
* `pmtu_discovery` - Flag indicating whether path MTU is enable or not for Vedge VPN interface type Feature Template
* `flow_control` - Value of bidirectional flow control settings for Vedge VPN interface type Feature Template
* `tcp_mss` - Maximum segment size (MSS) of TPC SYN packets passing through the router for VPN interface type Feature Template
* `speed` - Value of speed to use when the remote end of the connection does not support autonegotiation for VPN interface type Feature Template
* `clear_dont_fragment` - Flag indicating Don't Fragment (DF) bit in the IPv4 packet header for packets being transmitted out the interface is clar or not for VPN interface type Feature Template
* `static_ingess_qos` - Queue number to use for incoming traffic for Vedge VPN interface type Feature Template
* `arp_timeout` - Value indicating how long it takes for a dynamically learned ARP entry to time out for VPN interface type Feature Template
* `autonegotiation` - Flag indicating autonegotiation mode for VPN interface type Feature Template
* `tloc_extension` - Name of a physical interface on the same router that connects to the WAN transport for VPN interface type Feature Template
* `power_over_ethernet` - Flag indicating whether PoE(Power over Ethernet) is enabled or not for Vedge VPN interface type Feature Template
* `load_interval` - Value of load interval for Cisco VPN interface type Feature Template
* `tracker` - Name of a tracker to track the status of transport interfaces that connect to the internet for VPN interface type Feature Template
* `icmp_redirect_disable` - Flag indicating whether ICMP redirect disable or not on interface for VPN interface type Feature Template
* `gre_tunnel_source_ip` - IP address of the extended WAN interface
* `xconnect` - The name of a physical interface on the same router that connects to the WAN transport for VPN interface type Feature Template
* `ip_directed_broadcast` - Flag indicate whether directed IP braodcast is enabled or not for VPN interface  type Feature Template


<a id="nestedblock--ipv4"></a>

## Nested Schema for Basic ipv4
* `primary_address` - Value of IPv4 primary address for VPN interface  type Feature Template
* `secondary_address` - Value of IPv4 secondary address for VPN interface  type Feature Template(see [below for nested schema](#nestedblock--secondary_addressv4))
* `dhcp_distance` - Value of IPv4 DHCP distance for VPN interface type Feature Template
* `dhcp_helper` - Value of IPv4 DHCP Helper for VPN interface type Feature Template


<a id="nestedblock--ipv6"></a>

## Nested Schema for Basic ipv6
* `primary_address` - Value of IPv6 primary address for VPN interface type Feature Template
* `secondary_address` - Value of IPv6 secondary address for VPN interface type Feature Template(see [below for nested schema](#nestedblock--secondary_addressv6))
* `dhcp_distance` - Value of IPv6 DHCP distance for Vedge VPN interface type Feature Template
* `dhcp_helper` - Value of IPv6 DHCP Helper for VPN interface type Feature Template
* `dhcp_rapid_commit` - Value of IPv6 DHCP rapid commit enabled or not for Vedge VPN interface type Feature Template
* `dhcp_helper` - Value of IPv6 DHCP rapid commit enabled or not for VPN interface type Feature Template(see [below for nested schema](#nestedblock--dhcp_helper))


<a id="nestedblock--allow_service"></a>

## Nested Schema for Tunnel allow_service
* `all` - Flag indicating whether All serivce is enabled or not for VPN interface type Feature Template
* `bgp` - Flag indicating whether BGP is enabled or not for VPN interface type Feature Template
* `dhcp` - Flag indicating whether DHCP is enabled or not for VPN interface type Feature Template
* `dns` - Flag indicating whether DNS is enabled or not for VPN interface type Feature Template
* `icmp` - Flag indicating whether ICMP is enabled or not for VPN interface type Feature Template
* `netconf` - Flag indicating whether NETCONF is enabled or not for VPN interface type Feature Template
* `ntp` - Flag indicating whether NTP is enabled or not for VPN interface type Feature Template
* `ospf` - Flag indicating whether OSPF is enabled or not for VPN interface type Feature Template
* `ssh` - Flag indicating whether SSH is enabled or not for VPN interface type Feature Template
* `stun` - Flag indicating whether STUN is enabled or not for VPN interface type Feature Template
* `https` - Flag indicating whether HTTPS is enabled or not for VPN interface type Feature Template
* `snmp` - Flag indicating whether SNMP is enabled or not for Cisco VPN interface type Feature Template


<a id="nestedblock--encapsulation"></a>

## Nested Schema for Tunnel encapsulation
* `gre_preference` - Value of GRE preference for VPN interface type Feature Template
* `gre_weight` - Value of GRE weight for VPN interface type Feature Template
* `ipsec_preference` - Value of IPsec preference for VPN interface type Feature Template
* `ipsec_weight` - Value of IPsec weight for VPN interface type Feature Template
* `carrier` - Carrier name or private network identifier to associate with the tunnel for VPN interface type Feature Template
* `bind_loopback_tunnel` - Name of a physical interface to bind to a loopback interface for VPN interface type Feature Template
* `last_resort_circuit` - Flag indicating whether tunnel interface is used as the circuit of last resort for VPN interface type Feature Template
* `hold_time` - Value of hold time for Vedge VPN interface type Feature Template
* `nat_refresh_interval` - Value of interval between NAT refresh packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template
* `hello_interval` - Value of interval between Hello packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template
* `hello_tolerance` - Value of the time to wait for a Hello packet on a DTLS or TLS WAN transport connection before declaring that transport tunnel to be down for VPN interface type Feature Template
* `gre_tunnel_destination_ip` - Value of GRE tunnel destination IP for VPN interface type Feature Template



<a id="nestedblock--natv4"></a>

## Nested Schema for NAT ipv4
* `nat_type` - NAT type for Cisco VPN interface type Feature Template
* `refresh_mode` - Refresh mode for VPN interface type Feature Template
* `log_nat_flow_creations_or_deletions` - Flag indicating whether log nat flow creations or deletions are enabled or not for VPN interface type Feature Template
* `udp_timeout` - The time when NAT translations over UDP sessions time out for VPN interface type Feature Template
* `tcp_timeout` - The time when NAT translations over TCP sessions time out for VPN interface type Feature Template
* `block_icmp` - Flag indicating whether to block icmp or not for VPN interface type Feature Template
* `respond_to_ping` - Flag indicating whether to respond to ping or not for VPN interface type Feature Template
* `nat_pool_range_start` - Value of starting IP address for the NAT pool for VPN interface type Feature Template
* `nat_pool_range_end` - Value of closing IP address for the NAT pool for VPN interface type Feature Template
* `overload` - Flag indicating whether to enable per-port translation or not for Cisco VPN interface type Feature Template
* `nat_inside_source_loopback_interface` - Value of NAT inside source loopback interface for Cisco VPN interface type Feature Template
* `port_forward` - Configure port forwarding rule for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--port_forward))
* `static_nat` - Configure static nat for VPN interface type Feature Template(see [below for nested schema](#nestedblock--static_nat))


<a id="nestedblock--natv6"></a>

## Nested Schema for NAT ipv6
* `nat64` - Flag indicating whether to enable nat64 or not for VPN interface type Feature Template


<a id="nestedblock--vrrpv4"></a>

## Nested Schema for VRRP ipv4
* `group_id` - The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template
* `priority` - The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template
* `timer` - Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template
* `track_omp` - Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template
* `track_prefix_list` - (Required) Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template
* `ip_address` - (Optional) Value of IPv4 address of the virtual router for VPN interface type Feature Template


<a id="nestedblock--vrrpv6"></a>

## Nested Schema for VRRP ipv6
* `group_id` - The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template
* `priority` - The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template
* `timer` - Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template
* `track_omp` - Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template
* `track_prefix_list` - Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template
* `link_local_ipv6_address` - Value of IPv6 address of the virtual router for VPN interface type Feature Template
* `global_ipv6_prefix` - Value of global IPv6 prefix of the virtual router for VPN interface type Feature Template



<a id="nestedblock--advanced_options"></a>

## Nested Schema for 8021.X advanced_options
* `vlan` - Configure VLAN for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--vlan))
* `dynamic_authentication_server` - Configure Dynamic Authentication Server for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--dynamic_authentication_server))
* `mac_authentication_bypass` - Configure MAC Authentication Bypass for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--mac_authentication_bypass))
* `request_attributes` - Configure Request Attributes for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--request_attributes))


<a id="nestedblock--secondary_addressv4"></a>

## Nested Schema for Basic IPv4 secondary_address
* `address` - Value of IPv4 prefix for VPN interface type Feature Template


<a id="nestedblock--secondary_addressv6"></a>

## Nested Schema for Basic IPv6 secondary_address
* `address` - Value of IPv6 prefix for VPN interface type Feature Template


<a id="nestedblock--dhcp_helper"></a>

## Nested Schema for Basic IPv6 dhcp_helper
* `address` - Value of IPv6 address for VPN interface type Feature Template
* `vpn` - Value of VPN id for VPN interface type Feature Template



<a id="nestedblock--port_forward"></a>

## Nested Schema for NAT ipv4 port_forward
* `port_start_range` - Value of port starting range for Vedge VPN interface type Feature Template
* `port_end_range` - Value of port ending range for Vedge VPN interface type Feature Template
* `protocol` - Protocol for Vedge VPN interface type Feature Template
* `vpn` - VPN id for Vedge VPN interface type Feature Template
* `private_ip` - (Required) Private IP address for Vedge VPN interface type Feature Template


<a id="nestedblock--static_nat"></a>

## Nested Schema for NAT ipv4 static_nat
* `source_ip` - Value of the inside local address as source IP address for VPN interface type Feature Template
* `translated_source_ip` - Value of the inside global address as the translated source IP address for VPN interface type Feature Template
* `source_vpn_id` - Configures Source VPN ID, which is the Service VPN in which the host resides for VPN interface type Feature Template
* `static_nat_direction` - Set the direction in which to perform network address translation for VPN interface type Feature Template
* `protocol` - Protocol for Vedge VPN interface type Feature Template
* `source_port` - Value of source port for Vedge VPN interface type Feature Template
* `translate_port` - Value of translate port for Vedge VPN interface type Feature Template

<a id="nestedblock--vlan"></a>

## Nested Schema for vlan
* `authentication_fail_vlan` - Value of authentication fail VLAN for Vedge VPN interface type Feature Template
* `guest_vlan` - Value of guest VLAN for Vedge VPN interface type Feature Template
* `authentication_reject_vlan` - Value of authentication reject VLAN for Vedge VPN interface type Feature Template
* `default_vlan` - Value of default VLAN for Vedge VPN interface type Feature Template


<a id="nestedblock--dynamic_authentication_server"></a>

## Nested Schema for dynamic_authentication_server
* `das_port` - Value of Dynamic Authentication Server port for Vedge VPN interface type Feature Template
* `client` - Value of Client IP address for Vedge VPN interface type Feature Template
* `secret_key` - Secret key Value for Vedge VPN interface type Feature Template
* `time_window` - Time window value for Vedge VPN interface type Feature Template
* `required_timestamp` - Flag indicating whether to enable required timestamp or not for Vedge VPN interface type Feature Template
* `vpn` - Value of VPN ID for Vedge VPN interface type Feature Template


<a id="nestedblock--mac_authentication_bypass"></a>

## Nested Schema for mac_authentication_bypass
* `server` - Flag indicating whether to enable server or not for Vedge VPN interface type Feature Template
* `mac_authentication_bypass_entries` - List of MAC Authenticaion bypass entries for Vedge VPN interface type Feature Template

<a id="nestedblock--request_attributes"></a>

## Nested Schema for request_attributes
* `authentication` - Configure Authentication for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--authentication))
* `accounting` - Configure Accounting for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--accounting))

## Nested Schema for request_attributes authentication
* `id` - ID for Vedge VPN interface type Feature Template
* `syntax_choice` - Syntax Choice for Vedge VPN interface type Feature Template
* `value` - Value for Vedge VPN interface type Feature Template


<a id="nestedblock--accounting"></a>

## Nested Schema for request_attributes accounting
* `id` - ID for Vedge VPN interface type Feature Template
* `syntax_choice` - Syntax Choice for Vedge VPN interface type Feature Template
* `value` - Value for Vedge VPN interface type Feature Template














