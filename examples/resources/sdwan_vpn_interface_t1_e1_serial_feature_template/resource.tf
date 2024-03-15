resource "sdwan_vpn_interface_t1_e1_serial_feature_template" "example" {
  name                  = "Example"
  description           = "My Example"
  device_types          = ["vedge-C8000V"]
  serial_interface_name = "Example"
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
  encapsulation_for_tloc = [
    {
      encapsulation_type = "gre"
      preference         = 4294967
      weight             = 250
    }
  ]
  groups                                  = [42949672]
  border                                  = true
  per_tunnel_qos                          = true
  per_tunnel_qos_aggregator               = false
  color                                   = "custom1"
  last_resort_circuit                     = false
  low_bandwidth_link                      = false
  tunnel_tcp_mss                          = 1460
  enable_clear_don                        = false
  network_broadcast_1                     = false
  max_control_connections                 = 8
  vbond_as_stun_server                    = false
  exclude_controller_group_list           = [100]
  vmanage_connection_preference           = 5
  port_hop                                = false
  restrict                                = false
  carrier                                 = "carrier1"
  nat_refresh_interval                    = 15
  hello_interval                          = 1000
  hello_tolerance                         = 12
  bind_loopback_tunnel                    = "12"
  all                                     = false
  network_broadcast_2                     = false
  bgp                                     = false
  dhcp                                    = true
  dns                                     = true
  icmp                                    = true
  ssh                                     = false
  ntp                                     = false
  netconf                                 = false
  ospf                                    = false
  stun                                    = false
  snmp                                    = false
  https                                   = true
  clear_dont_fragment_bit                 = false
  pmtu_discovery                          = false
  ip_mtu                                  = 1500
  static_ingress_qos                      = 6
  tcp_mss                                 = 720
  tloc_extension                          = "tloc"
  administrative_shutdown                 = true
  link_autonegotiate                      = true
  shaping_rate                            = 10000000
  qos_map                                 = "test"
  vpn_qos_map                             = "test"
  interface_bandwidth_capacity            = 128
  clock_rate                              = "5300000"
  encapsulation                           = "hdlc"
  interface_downstream_bandwidth_capacity = 10000000
  write_rule                              = "test_rule"
  access_list = [
    {
      direction = "in"
      acl_name  = "Egress ACL - IPv4"
    }
  ]
}
