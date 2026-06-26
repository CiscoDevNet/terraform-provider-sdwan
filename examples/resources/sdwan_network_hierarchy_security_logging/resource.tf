resource "sdwan_network_hierarchy_security_logging" "example" {
  high_speed_logging = [
    {
      vrf       = "service_lan_vpn1"
      server_ip = "10.1.2.1"
      port      = 2055
    }
  ]
  utd_syslog = [
    {
      vpn       = "service_lan_vpn1"
      server_ip = "10.1.1.2"
    }
  ]
}
