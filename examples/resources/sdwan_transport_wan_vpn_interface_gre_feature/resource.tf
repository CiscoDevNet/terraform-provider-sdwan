resource "sdwan_transport_wan_vpn_interface_gre_feature" "example" {
  name                            = "Example"
  description                     = "My Example"
  feature_profile_id              = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  transport_wan_vpn_feature_id    = "140331f6-5418-4755-a059-13c77eb96037"
  interface_name                  = "gre1"
  interface_description           = "gre1"
  ipv4_address                    = "70.1.1.1"
  ipv4_subnet_mask                = "255.255.255.0"
  shutdown                        = true
  tunnel_source_ipv4_address      = "78.1.1.1"
  tunnel_destination_ipv4_address = "79.1.1.1"
  ip_mtu                          = 1500
  tcp_mss                         = 1460
  clear_dont_fragment             = false
  application_tunnel_type         = "none"
}
