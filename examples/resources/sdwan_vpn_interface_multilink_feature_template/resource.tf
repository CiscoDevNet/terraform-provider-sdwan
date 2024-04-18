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
      acl_name  = "ACL1"
    }
  ]
  ppp_authentication_protocol     = "chap"
  ppp_authentication_protocol_pap = false
  chap_hostname                   = "chap-example"
  chap_ppp_auth_password          = "myPassword"
  pap_username                    = "pap-username"
  pap_password                    = true
  pap_ppp_auth_password           = "myPassword"
  ppp_authentication_type         = "callin"
  enable_core_region              = true
  core_region                     = "core"
  secondary_region                = "off"
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
  disable_fragmentation                          = true
  fragment_max_delay                             = 1
  interleaving_fragment                          = false
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
  bandwidth_upstream                             = 214748300
  bandwidth_downstream                           = 214748300
  write_rule                                     = "RULE1"
  access_lists = [
    {
      direction = "in"
      acl_name  = "ACL2"
    }
  ]
  multilink_interfaces = [
    {
      interface_type       = "E1"
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
          time_slot     = ["example"]
        }
      ]
    }
  ]
  nim_interface_list = [
    {
      nim_serial_interface_type = "2T"
      interface_name            = "nim-interface"
      interface_description     = "My Description"
      bandwidth                 = 21474836
      clock_rate                = 120000
      encapsulation_serial      = "hdlc"
    }
  ]
}
