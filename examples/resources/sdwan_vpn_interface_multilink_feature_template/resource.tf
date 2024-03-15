resource "sdwan_vpn_interface_multilink_feature_template" "example" {
  name                   = "Example"
  description            = "My Example"
  device_types           = ["vedge-C8000V"]
  interface_name         = "Example"
  multilink_group_number = 2147483
  interface_description  = "My Description"
  ipv4_address           = "1.2.3.4"
  ipv6_address           = "2001:0:0:1::/64"
  ipv6_access_lists = [
    {
      direction = "in"
      acl_name  = "Egress ACL - IPv4"
    }
  ]
  ppp_authentication_protocol     = "chap"
  ppp_authentication_protocol_pap = false
  authentication_type             = "callin"
  chap_hostname                   = "chap-example"
  chap_ppp_auth_password          = "myPassword"
  pap_username                    = "pap-username"
  pap_password                    = true
  pap_ppp_auth_password           = "myPassword"
  enable_core_region              = true
  core_region                     = "core"
  secondary_region                = "off"
  encapsulation = [
    {
      encapsulation_type = "gre"
      preference         = 4294967
      weight             = 250
    }
  ]
  groups                        = [42949672]
  border                        = true
  per_tunnel_qos                = true
  per_tunnel_qos_aggregator     = false
  color                         = "custom1"
  last_resort_circuit           = false
  low_bandwidth_link            = false
  tunnel_tcp_mss                = 1460
  enable_clear_dont_fragment    = false
  network_broadcast_1           = false
  max_control_connections       = 8
  control_connections           = true
  vbond_as_stun_server          = false
  exclude_controller_group_list = [100]
  vmanage_connection_preference = 5
  port_hop                      = false
  restrict                      = false
  carrier                       = "carrier1"
  nat_refresh_interval          = 15
  hello_interval                = 1000
  hello_tolerance               = 12
  bind_loopback_tunnel          = "12"
  all                           = false
  network_broadcast_2           = false
  bgp                           = false
  dhcp                          = true
  dns                           = true
  icmp                          = true
  ssh                           = false
  netconf                       = false
  ospf                          = false
  stun                          = false
  snmp                          = false
  https                         = true
  disable_fragmentation         = true
  fragment_max_delay            = 1
  interleaving_fragment         = false
  clear_dont_fragment_bit       = false
  pmtu_discovery                = false
  ip_mtu                        = 1500
  static_ingress_qos            = 6
  tcp_mss                       = 720
  ip_directed_broadcast         = true
  tloc_extension                = "tloc"
  administrative_shutdown       = true
  link_autonegotiate            = true
  shaping_rate                  = 10000000
  qos_map                       = "test"
  vpn_qos_map                   = "test"
  bandwidth_upstream            = 214748300
  bandwidth_downstream          = 214748300
  write_rule                    = "test_rule"
  access_list = [
    {
      direction = "in"
      acl_name  = "Egress ACL - IPv4"
    }
  ]
  controller_tx_ex_list = [
    {
      card_type            = "E1"
      slot                 = "interface-t1"
      framing              = "example-framing"
      line_mode            = "primary"
      internal             = false
      description          = "example-interface"
      linecode             = "ami"
      set_length_for_long  = "100"
      set_length_for_short = "100"
      channel_group_list = [
        {
          channel_group = 30
          timeslots     = ["example"]
        }
      ]
    }
  ]
  nim_interface_list = [
    {
      nim_serial_interface_type = "2t"
      interface_name            = "nim-interface"
      interface_description     = "interface description"
      bandwidth                 = 21474836
      clock_rate                = 120000
      encapsulation_serial      = "hdlc"
    }
  ]
}
