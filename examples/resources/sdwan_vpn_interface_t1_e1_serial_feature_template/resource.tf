resource "sdwan_vpn_interface_t1_e1_serial_feature_template" "example" {
  name                  = "Example"
  description           = "My Example"
  device_types          = ["vedge-C8000V"]
  serial_interface_name = "SERIAL1"
  interface_description = "My description"
  ipv4_address          = "1.2.3.4/24"
  ipv6_address          = "2001:0:0:1::/64"
  ipv6_access_lists = [
    {
      direction = "in"
      acl_name  = "ACL1"
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
  tunnel_interface_control_connections           = 8
  tunnel_interface_vbond_as_stun_server          = false
  tunnel_interface_exclude_controller_group_list = [100]
  tunnel_interface_vmanage_connection_preference = 5
  tunnel_interface_port_hop                      = false
  tunnel_interface_restrict                      = false
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
  shutdown                                       = true
  autonegotiate                                  = true
  shaping_rate                                   = 10000000
  qos_map                                        = "test"
  qos_map_vpn                                    = "test"
  interface_bandwidth_capacity                   = 128
  clock_rate                                     = "5300000"
  encapsulation                                  = "hdlc"
  interface_downstream_bandwidth_capacity        = 10000000
  write_rule                                     = "RULE1"
  access_lists = [
    {
      direction = "in"
      acl_name  = "ACL2"
    }
  ]
}
