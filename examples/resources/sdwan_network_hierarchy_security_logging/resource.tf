resource "sdwan_network_hierarchy_security_logging" "example" {
  node_id = "5333c142-394b-47a7-afa6-760c44ca3cb5"
  high_speed_logging = [
    {
      vrf       = "1"
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
