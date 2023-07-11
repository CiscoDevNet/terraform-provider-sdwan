resource "sdwan_cisco_dhcp_server_feature_template" "example" {
  name              = "Example"
  description       = "My Example"
  device_types      = ["vedge-C8000V"]
  address_pool      = "10.1.1.0/24"
  exclude_addresses = ["10.1.1.1-10.1.1.5", "10.1.1.254"]
  lease_time        = 600
  interface_mtu     = 1500
  domain_name       = "cisco.com"
  default_gateway   = "10.1.1.254"
  dns_servers       = ["1.2.3.4"]
  tftp_servers      = ["1.2.3.4"]
  static_leases = [
    {
      mac_address = "11:11:11:11:11:11"
      ip_address  = "10.1.1.10"
      hostname    = "HOST1"
    }
  ]
  options = [
    {
      option_code = 10
      ascii       = "abc"
    }
  ]
}
