resource "sdwan_system_global_profile_parcel" "example" {
  name                                       = "Example"
  description                                = "My Example"
  feature_profile_id                         = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  services_global_services_ip_http_server    = false
  services_global_services_ip_https_server   = false
  services_global_services_ip_ftp_passive    = false
  services_global_services_ip_domain_lookup  = false
  services_global_services_ip_arp_proxy      = false
  services_global_services_ip_rcmd           = false
  services_global_services_ip_line_vty       = false
  services_global_services_ip_cdp            = true
  services_global_services_ip_lldp           = true
  services_global_services_ip_source_intrf   = "GigabitEthernet0/0/1"
  global_other_settings_tcp_keepalives_in    = true
  global_other_settings_tcp_keepalives_out   = true
  global_other_settings_tcp_small_servers    = false
  global_other_settings_udp_small_servers    = false
  global_other_settings_console_logging      = true
  global_other_settings_i_p_source_route     = false
  global_other_settings_vty_line_logging     = false
  global_other_settings_snmp_ifindex_persist = true
  global_other_settings_ignore_bootp         = true
  global_settings_nat64_udp_timeout          = 300
  global_settings_nat64_tcp_timeout          = 3600
  global_settings_http_authentication        = "aaa"
  global_settings_s_s_h_version              = "2"
}
