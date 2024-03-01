resource "sdwan_cisco_vpn_interface_feature_template" "example" {
  name                  = "Example"
  description           = "My Example"
  device_types          = ["vedge-C8000V"]
  interface_name        = "ge0/0"
  interface_description = "My Interface Description"
  poe                   = false
  address               = "1.1.1.1/24"
  ipv4_secondary_addresses = [
    {
      address = "2.2.2.2/24"
    }
  ]
  dhcp          = false
  dhcp_distance = 10
  ipv6_address  = "2001:1::1/48"
  dhcpv6        = false
  ipv6_secondary_addresses = [
    {
      address = "2.2.2.2/24"
    }
  ]
  ipv6_access_lists = [
    {
      direction = "in"
      acl_name  = "ACL1"
    }
  ]
  ipv4_dhcp_helper = ["6.6.6.6"]
  ipv6_dhcp_helpers = [
    {
      address = "2001:7::7/48"
      vpn_id  = 5
    }
  ]
  tracker                              = ["tracker1"]
  auto_bandwidth_detect                = false
  iperf_server                         = "8.8.8.8"
  nat                                  = true
  nat_type                             = "interface"
  udp_timeout                          = 1
  tcp_timeout                          = 60
  nat_pool_range_start                 = "10.1.1.1"
  nat_pool_range_end                   = "10.1.1.255"
  nat_overload                         = false
  nat_inside_source_loopback_interface = "lo1"
  nat_pool_prefix_length               = 24
  ipv6_nat                             = false
  nat64_interface                      = false
  nat66_interface                      = false
  static_nat66_entries = [
    {
      source_prefix            = "2001:7::/48"
      translated_source_prefix = "2001:8::/48"
      source_vpn_id            = 1
    }
  ]
  static_nat_entries = [
    {
      source_ip            = "10.1.1.1"
      translate_ip         = "100.1.1.1"
      static_nat_direction = "inside"
      source_vpn_id        = 1
    }
  ]
  static_port_forward_entries = [
    {
      source_ip            = "10.1.1.1"
      translate_ip         = "100.1.1.1"
      static_nat_direction = "inside"
      source_port          = 8000
      translate_port       = 9000
      protocol             = "tcp"
      source_vpn_id        = 1
    }
  ]
  enable_core_region = false
  core_region        = "core"
  secondary_region   = "off"
  tunnel_interface_encapsulations = [
    {
      encapsulation = "gre"
      preference    = 10
      weight        = 100
    }
  ]
  tunnel_interface_border                        = false
  tunnel_qos_mode                                = "spoke"
  tunnel_bandwidth                               = 50
  tunnel_interface_groups                        = [5]
  tunnel_interface_color                         = "gold"
  tunnel_interface_max_control_connections       = 10
  tunnel_interface_control_connections           = false
  tunnel_interface_vbond_as_stun_server          = false
  tunnel_interface_exclude_controller_group_list = [10]
  tunnel_interface_vmanage_connection_preference = 5
  tunnel_interface_port_hop                      = false
  tunnel_interface_color_restrict                = false
  tunnel_interface_gre_tunnel_destination_ip     = "5.5.5.5"
  tunnel_interface_carrier                       = "carrier1"
  tunnel_interface_nat_refresh_interval          = 5
  tunnel_interface_hello_interval                = 1000
  tunnel_interface_hello_tolerance               = 12
  tunnel_interface_bind_loopback_tunnel          = "1"
  tunnel_interface_last_resort_circuit           = false
  tunnel_interface_low_bandwidth_link            = false
  tunnel_interface_tunnel_tcp_mss                = 1460
  tunnel_interface_clear_dont_fragment           = false
  tunnel_interface_propagate_sgt                 = false
  tunnel_interface_network_broadcast             = false
  tunnel_interface_allow_all                     = false
  tunnel_interface_allow_bgp                     = false
  tunnel_interface_allow_dhcp                    = false
  tunnel_interface_allow_dns                     = false
  tunnel_interface_allow_icmp                    = false
  tunnel_interface_allow_ssh                     = false
  tunnel_interface_allow_netconf                 = false
  tunnel_interface_allow_ntp                     = false
  tunnel_interface_allow_ospf                    = false
  tunnel_interface_allow_stun                    = false
  tunnel_interface_allow_snmp                    = false
  tunnel_interface_allow_https                   = false
  media_type                                     = "auto-select"
  interface_mtu                                  = 9216
  ip_mtu                                         = 1500
  tcp_mss_adjust                                 = 1460
  tloc_extension                                 = "123"
  load_interval                                  = 30
  gre_tunnel_source_ip                           = "3.3.3.3"
  gre_tunnel_xconnect                            = "a123"
  mac_address                                    = "00-B0-D0-63-C2-26"
  speed                                          = "1000"
  duplex                                         = "full"
  shutdown                                       = false
  arp_timeout                                    = 1200
  autonegotiate                                  = true
  ip_directed_broadcast                          = false
  icmp_redirect_disable                          = false
  qos_adaptive_period                            = 15
  qos_adaptive_bandwidth_downstream              = 10000
  qos_adaptive_min_downstream                    = 100
  qos_adaptive_max_downstream                    = 100000
  qos_adaptive_bandwidth_upstream                = 10000
  qos_adaptive_min_upstream                      = 100
  qos_adaptive_max_upstream                      = 100000
  shaping_rate                                   = 1000
  qos_map                                        = "QOSMAP1"
  qos_map_vpn                                    = "QOSMAP2"
  bandwidth_upstream                             = 10000
  bandwidth_downstream                           = 10000
  block_non_source_ip                            = false
  rewrite_rule_name                              = "RULE1"
  access_lists = [
    {
      direction = "in"
      acl_name  = "ACL1"
    }
  ]
  static_arps = [
    {
      ip_address = "8.8.8.8"
      mac        = "00-B0-D0-63-C2-26"
    }
  ]
  ipv4_vrrps = [
    {
      group_id          = 100
      priority          = 100
      timer             = 100
      track_omp         = false
      track_prefix_list = "PL1"
      ip_address        = "2.2.2.2"
      ipv4_secondary_addresses = [
        {
          ip_address = "2.2.2.3"
        }
      ]
      tloc_preference_change       = false
      tloc_preference_change_value = 10
      tracking_objects = [
        {
          tracker_id      = 10
          track_action    = "decrement"
          decrement_value = 100
        }
      ]
    }
  ]
  ipv6_vrrps = [
    {
      group_id          = 100
      priority          = 100
      timer             = 100
      track_omp         = false
      track_prefix_list = "PL1"
      ipv6_addresses = [
        {
          ipv6_link_local = "fe80::260:8ff:fe52:f9d8"
          prefix          = "2001:9::/48"
        }
      ]
    }
  ]
  propagate_sgt       = false
  static_sgt          = 1003
  static_sgt_trusted  = false
  enable_sgt          = true
  sgt_enforcement     = true
  sgt_enforcement_sgt = 1004
}
