resource "sdwan_cisco_vpn_interface_gre_feature_template" "example" {
  name                    = "Example"
  description             = "My Example"
  device_types            = ["vedge-C8000V"]
  interface_name          = "gre0/0"
  interface_description   = "My Description"
  ip_address              = "1.1.1.1/24"
  tunnel_source           = "1.2.3.4"
  shutdown                = true
  tunnel_source_interface = "e1"
  tunnel_destination      = "3.4.5.6"
  application             = "sig"
  ip_mtu                  = 1500
  tcp_mss_adjust          = 1400
  clear_dont_fragment     = true
  rule_name               = "test"
  access_list = [
    {
      direction = "in"
      acl_name  = "test"
    }
  ]
  tracker          = ["TRACKER1"]
  tunnel_route_via = "g0/0"
}
