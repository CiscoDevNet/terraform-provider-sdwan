resource "sdwan_service_switchport_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  interfaces = [
    {
      interface_name                 = "GigabitEthernet"
      mode                           = "access"
      shutdown                       = true
      speed                          = "10"
      duplex                         = "full"
      switchport_access_vlan         = 1
      switchport_trunk_allowed_vlans = "1"
      switchport_trunk_native_vlan   = 1
      enable_dot1x                   = false
      port_control                   = "auto"
      voice_vlan                     = 1
      pae_enable                     = true
      mac_authentication_bypass      = false
      host_mode                      = "single-host"
      enable_periodic_reauth         = false
      inactivity                     = 60
      reauthentication               = 1
      control_direction              = "both"
      restricted_vlan                = 1
      guest_vlan                     = 1
      critical_vlan                  = 1
      enable_voice                   = false
    }
  ]
  age_out_time = 300
  static_mac_addresses = [
    {
      mac_address    = "01:02:03:04:05:06"
      vlan_id        = 1
      interface_name = "GigabitEthernet0/0/0"
    }
  ]
}
