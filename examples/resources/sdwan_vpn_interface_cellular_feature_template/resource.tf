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
