resource "sdwan_vpn_interface_cellular_feature_template" "example" {
  name                    = "Example"
  description             = "My Example"
  device_types            = ["vedge-C8000V"]
  cellular_interface_name = "Example"
  interface_description   = "My Description"
  ipv6_access_lists = [
    {
      direction = "in"
      acl_name  = "Egress ACL - IPv6"
    }
  ]
  ipv4_dhcp_helper = ["6"]
  tracker          = ["tracker1"]
  nat              = true
  refresh_mode     = "outbound"
  udp_timeout      = 1
  tcp_timeout      = 60
  block_icmp_error = true
  respond_to_ping  = false
  port_forward = [
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
  encapsulation = [
    {
      encapsulation_type = "gre"
      preference         = 4294967
      weight             = 250
    }
  ]
  groups                          = [42949672]
  border                          = true
  per_tunnel_qos                  = true
  per_tunnel_qos_aggregator       = false
  color                           = "custom1"
  last_resort_circuit             = false
  low_bandwidth_link              = false
  tunnel_tcp_mss                  = 1460
  enable_clear_dont_fragment      = false
  network_broadcast_1             = false
  max_control_connections         = 8
  control_connections             = true
  vbond_as_stun_server            = false
  exclude_controller_group_list   = [100]
  vmanage_connection_preference   = 5
  port_hop                        = false
  restrict                        = false
  carrier                         = "carrier1"
  nat_refresh_interval            = 15
  hello_interval                  = 1000
  hello_tolerance                 = 12
  bind_loopback_tunnel            = "12"
  all                             = false
  network_broadcast_2             = false
  bgp                             = false
  dhcp                            = true
  dns                             = true
  icmp                            = true
  ssh                             = false
  ntp                             = false
  netconf                         = false
  ospf                            = false
  stun                            = false
  snmp                            = false
  https                           = true
  clear_dont_fragment_bit         = false
  pmtu_discovery                  = false
  ip_mtu                          = 1500
  static_ingress_qos              = 6
  tcp_mss                         = 720
  tloc_extension                  = "tloc"
  ip_directed_broadcast           = true
  administrative_shutdown         = true
  link_autonegotiate              = true
  adaptive_qos                    = true
  adapt_period                    = 15
  adaptive_qos_downstream_default = 10000000
  downstream_min_bandwidth_limit  = 10000000
  downstream_max_bandwidth_limit  = 10000000
  adaptive_qos_upstream_default   = 10000000
  upstream_min_bandwidth_limit    = 10000000
  upstream_max_bandwidth_limit    = 10000000
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
  static_arp_entries = [
    {
      ip_address  = "1.2.3.4"
      mac_address = "00-B0-D0-63-C2-26"
    }
  ]
}
