resource "sdwan_transport_wan_vpn_interface_ethernet_feature" "example" {
  name                         = "Example"
  description                  = "My Example"
  feature_profile_id           = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  transport_wan_vpn_feature_id = "140331f6-5418-4755-a059-13c77eb96037"
  shutdown                     = true
  interface_name               = "GigabitEthernet1"
  interface_description        = "WAN"
  ipv4_configuration_type      = "static"
  ipv4_address                 = "1.2.3.4"
  ipv4_subnet_mask             = "0.0.0.0"
  ipv4_secondary_addresses = [
    {
      address     = "1.2.3.4"
      subnet_mask = "0.0.0.0"
    }
  ]
  ipv4_dhcp_helper                               = ["1.2.3.4"]
  ipv6_configuration_type                        = "static"
  ipv6_address                                   = "2001:0:0:1::1/64"
  iperf_server                                   = "example"
  block_non_source_ip                            = false
  service_provider                               = "example"
  bandwidth_upstream                             = 21474836
  bandwidth_downstream                           = 21474836
  auto_detect_bandwidth                          = false
  tunnel_interface                               = true
  per_tunnel_qos                                 = true
  tunnel_qos_mode                                = "hub"
  tunnel_bandwidth_percent                       = 82
  tunnel_interface_bind_loopback_tunnel          = "example"
  tunnel_interface_carrier                       = "default"
  tunnel_interface_color                         = "mpls"
  tunnel_interface_hello_interval                = 1000
  tunnel_interface_hello_tolerance               = 12
  tunnel_interface_last_resort_circuit           = false
  tunnel_interface_gre_tunnel_destination_ip     = "1.2.3.4"
  tunnel_interface_color_restrict                = true
  tunnel_interface_groups                        = 42949672
  tunnel_interface_border                        = false
  tunnel_interface_max_control_connections       = 62
  tunnel_interface_nat_refresh_interval          = 5
  tunnel_interface_vbond_as_stun_server          = false
  tunnel_interface_exclude_controller_group_list = [2]
  tunnel_interface_vmanage_connection_preference = 8
  tunnel_interface_port_hop                      = true
  tunnel_interface_low_bandwidth_link            = false
  tunnel_interface_tunnel_tcp_mss                = 1460
  tunnel_interface_clear_dont_fragment           = false
  tunnel_interface_cts_sgt_propagation           = false
  tunnel_interface_network_broadcast             = false
  tunnel_interface_allow_fragmentation           = false
  tunnel_interface_set_sdwan_tunnel_mtu_to_max   = false
  tunnel_interface_allow_all                     = false
  tunnel_interface_allow_bgp                     = false
  tunnel_interface_allow_dhcp                    = true
  tunnel_interface_allow_ntp                     = false
  tunnel_interface_allow_ssh                     = false
  tunnel_interface_allow_dns                     = true
  tunnel_interface_allow_icmp                    = true
  tunnel_interface_allow_https                   = true
  tunnel_interface_allow_ospf                    = false
  tunnel_interface_allow_stun                    = false
  tunnel_interface_allow_snmp                    = false
  tunnel_interface_allow_netconf                 = false
  tunnel_interface_allow_bfd                     = false
  tunnel_interface_encapsulations = [
    {
      encapsulation = "gre"
      preference    = 4294967
      weight        = 250
    }
  ]
  nat_ipv4 = true
  nat_type = "interface"
  nat_ipv4_pools = [
    {
      name                          = 10
      range_start                   = "203.0.115.50"
      range_end                     = "203.0.115.100"
      overload                      = true
      prefix_length                 = 25
      enable_dual_router_ha_mapping = false
    }
  ]
  nat_ipv4_loopbacks = [
    {
      loopback_interface = "Loopback0"
    }
  ]
  nat_udp_timeout = 1
  nat_tcp_timeout = 60
  new_static_nats = [
    {
      source_ip                     = "1.2.3.4"
      translated_ip                 = "2.3.4.5"
      direction                     = "inside"
      source_vpn                    = 3
      enable_dual_router_ha_mapping = false
    }
  ]
  static_port_forwards = [
    {
      protocol                      = "tcp"
      source_ip                     = "1.2.3.4"
      source_port                   = 8080
      translated_ip                 = "2.3.4.5"
      translated_port               = 80
      direction                     = "inside"
      source_vpn                    = 3
      enable_dual_router_ha_mapping = false
    }
  ]
  nat_ipv6 = true
  nat64    = false
  nat66    = true
  static_nat66 = [
    {
      source_prefix            = "2001:0db8:85a3::/48"
      translated_source_prefix = "abcd:1234:5678::/48"
      source_vpn_id            = 4
      egress_interface         = true
    }
  ]
  qos_adaptive                = false
  qos_shaping_rate            = 16
  acl_ipv4_egress_feature_id  = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  acl_ipv6_ingress_feature_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  arps = [
    {
      ip_address  = "1.2.3.4"
      mac_address = "00-B0-D0-63-C2-26"
    }
  ]
  icmp_redirect_disable = true
  duplex                = "full"
  mac_address           = "00-B0-D0-63-C2-26"
  ip_mtu                = 1500
  interface_mtu         = 1500
  tcp_mss               = 505
  speed                 = "2500"
  arp_timeout           = 1200
  autonegotiate         = false
  media_type            = "rj45"
  tloc_extension        = "tloc"
  gre_tunnel_source_ip  = "1.2.3.4"
  xconnect              = "example"
  load_interval         = 30
  tracker               = "example"
  ip_directed_broadcast = false
}
