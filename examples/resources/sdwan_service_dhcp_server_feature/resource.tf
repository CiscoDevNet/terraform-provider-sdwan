resource "sdwan_service_dhcp_server_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  network_address    = "1.2.3.4"
  subnet_mask        = "255.255.255.0"
  exclude            = ["192.168.1.1"]
  lease_time         = 86400
  interface_mtu      = 65535
  domain_name        = "example.com"
  default_gateway    = "1.2.3.4"
  dns_servers        = ["8.8.8.8"]
  tftp_servers       = ["1.1.1.1"]
  static_leases = [
    {
      mac_address = "01:02:03:04:05:06"
      ip_address  = "1.2.3.4"
    }
  ]
  option_codes = [
    {
      code  = 250
      ascii = "example"
    }
  ]
}
