resource "sdwan_vpn_interface_dsl_ipoe_feature_template" "example" {
  name                     = "Example"
  description              = "My Example"
  device_types             = ["vedge-C8000V"]
  ethernet_interface_name  = "Example"
  ethernet_ipv4_address    = "1.2.3.4/24"
  ethernet_enable_dhcp     = false
  ethernet_dhcp_distance   = 1234
  ethernet_dhcp_helper     = ["3"]
  internal_controller_type = "ipoe"
  shutdown                 = true
  ethernet_description     = "My Description"
  vdsl_configurations = [
    {
      controller_vdsl_slot     = "Example"
      sra                      = true
      mode_adsl1               = false
      mode_adsl2               = false
      mode_adsl2plus           = false
      mode_vdsl2               = false
      mode_ansi                = false
      vdsl_modem_configuration = "100"
    }
  ]
  encap                           = 4094
  dialer_pool_number              = 255
  ppp_maximum_payload             = 1790
  dialer_address_negotiated       = false
  unnumbered_loopback_interface   = "example"
  ppp_authentication_protocol     = "chap"
  ppp_authentication_protocol_pap = false
  chap_hostname                   = "chap-example"
  chap_ppp_auth_password          = "myPassword"
  pap_username                    = "pap-username"
  pap_password                    = true
  pap_ppp_auth_password           = "myPassword"
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
  nat                                            = true
  nat_refresh_mode                               = "outbound"
  nat_udp_timeout                                = 1
  nat_tcp_timeout                                = 60
  nat_block_icmp_error                           = true
  nat_response_to_ping                           = false
  nat_port_forwards = [
    {
      port_start_range   = 0
      port_end_range     = 65530
      protocol           = "tcp"
      private_vpn        = 65530
      private_ip_address = "1.2.3.4"
    }
  ]
  qos_adaptive_period               = 15
  qos_adaptive_bandwidth_downstream = 10000
  qos_adaptive_min_downstream       = 100
  qos_adaptive_max_downstream       = 100000
  qos_adaptive_bandwidth_upstream   = 10000
  qos_adaptive_min_upstream         = 100
  qos_adaptive_max_upstream         = 100000
  shaping_rate                      = 10000000
  qos_map                           = "test"
  qos_map_vpn                       = "test"
  bandwidth_upstream                = 214748300
  bandwidth_downstream              = 214748300
  write_rule                        = "RULE1"
  access_lists = [
    {
      direction = "in"
      acl_name  = "ACL1"
    }
  ]
  policers = [
    {
      direction    = "in"
      policer_name = "example"
    }
  ]
  ip_mtu                = 1500
  tcp_mss               = 720
  tloc_extension        = "tloc"
  tracker               = ["tracker1"]
  ip_directed_broadcast = true
}
