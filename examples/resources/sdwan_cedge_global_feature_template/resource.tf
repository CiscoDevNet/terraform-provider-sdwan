resource "sdwan_cedge_global_feature_template" "example" {
  name                 = "Example"
  description          = "My Example"
  device_types         = ["vedge-C8000V"]
  nat64_udp_timeout    = 300
  nat64_tcp_timeout    = 3600
  http_authentication  = "local"
  ssh_version          = 2
  http_server          = true
  https_server         = true
  source_interface     = "e1"
  ip_source_routing    = true
  arp_proxy            = true
  ftp_passive          = true
  rsh_rcp              = true
  bootp                = true
  domain_lookup        = true
  tcp_keepalives_out   = true
  tcp_keepalives_in    = true
  tcp_small_servers    = true
  udp_small_servers    = true
  lldp                 = true
  cdp                  = true
  snmp_ifindex_persist = true
  console_logging      = true
  vty_logging          = true
  line_vty             = true
}
