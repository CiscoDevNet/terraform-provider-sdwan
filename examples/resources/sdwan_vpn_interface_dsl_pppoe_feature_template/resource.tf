resource "sdwan_vpn_interface_dsl_pppoe_feature_template" "example" {
  name                    = "Example"
  description             = "My Example"
  device_types            = ["vedge-C8000V"]
  ethernet_interface_name = "Example"
  shutdown                = true
  interface_description   = "My Description"
  vdsl = [
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
  authentication_type             = "callin"
  chap_hostname                   = "chap-example"
  chap_password                   = "myPassword"
  pap_username                    = "pap-username"
  pap_password                    = true
  pap_outbound_password           = "myPassword"
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
  clear_dont_fragment           = false
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
  nat                           = true
  refresh_mode                  = "outbound"
  udp_timeout                   = 1
  tcp_timeout                   = 60
  block_icmp_error              = true
  respond_to_ping               = false
  port_forward = [
    {
      port_start_range   = 0
      port_end_range     = 65530
      protocol           = "tcp"
      private_vpn        = 65530
      private_ip_address = "1.2.3.4"
    }
  ]
  adaptive_qos                    = true
  adapt_period                    = 15
  shaping_rate_downstream_default = 10000000
  shaping_rate_downstream_min     = 10000000
  shaping_rate_downstream_max     = 10000000
  shaping_rate_upstream_default   = 10000000
  shaping_rate_upstream_min       = 10000000
  shaping_rate_upstream_max       = 10000000
  shaping_rate                    = 10000000
  qos_map                         = "test"
  vpn_qos_map                     = "test"
  bandwidth_upstream              = 214748300
  bandwidth_downstream            = 214748300
  write_rule                      = "test_rule"
  access_list = [
    {
      direction = "in"
      acl_name  = "Egress ACL - IPv4"
    }
  ]
  policer = [
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
