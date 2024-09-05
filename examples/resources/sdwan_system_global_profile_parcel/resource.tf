resource "sdwan_system_global_feature" "example" {
  name                 = "Example"
  description          = "My Example"
  feature_profile_id   = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  http_server          = false
  https_server         = false
  ftp_passive          = false
  domain_lookup        = false
  arp_proxy            = false
  rsh_rcp              = false
  line_vty             = false
  cdp                  = true
  lldp                 = true
  source_interface     = "GigabitEthernet0/0/1"
  tcp_keepalives_in    = true
  tcp_keepalives_out   = true
  tcp_small_servers    = false
  udp_small_servers    = false
  console_logging      = true
  ip_source_routing    = false
  vty_line_logging     = false
  snmp_ifindex_persist = true
  ignore_bootp         = true
  nat64_udp_timeout    = 300
  nat64_tcp_timeout    = 3600
  http_authentication  = "aaa"
  ssh_version          = "2"
}
