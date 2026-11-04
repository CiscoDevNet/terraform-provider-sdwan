resource "sdwan_transport_wan_vpn_interface_dsl_ipoe_feature" "example" {
  name                         = "Example"
  description                  = "My Example"
  feature_profile_id           = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  transport_wan_vpn_feature_id = "140331f6-5418-4755-a059-13c77eb96037"
  shutdown                     = true
  vdsl_slot                    = "0/0/0"
  interface_name               = "GigabitEthernet1"
  ipv4_address_type            = "dynamic"
  ipv4_address                 = "192.0.2.10"
  ipv4_subnet_mask             = "255.255.255.0"
  tunnel_interface_enabled     = true
  tunnel_interface_color       = "default"
}
