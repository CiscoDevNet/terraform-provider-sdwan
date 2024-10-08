resource "sdwan_other_ucse_feature" "example" {
  name                             = "Example"
  description                      = "My Example"
  feature_profile_id               = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  bay                              = 2
  slot                             = 0
  access_port_dedicated            = false
  access_port_shared_type          = "ge1"
  access_port_shared_failover_type = "ge2"
  ipv4_address                     = "2.2.2.2/24"
  default_gateway                  = "2.2.2.2"
  vlan_id                          = 3
  assign_priority                  = 3
  interfaces = [
    {
      interface_name     = "ucse2/0"
      ucse_interface_vpn = 2
      ipv4_address       = "10.1.15.15/24"
    }
  ]
}
