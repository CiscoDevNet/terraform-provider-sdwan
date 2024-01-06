resource "sdwan_switchport_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  slot         = 0
  sub_slot     = 0
  module_type  = "4"
  interfaces = [
    {
      name                                     = "GigabitEthernet0/0/0"
      switchport_mode                          = "access"
      shutdown                                 = true
      speed                                    = "100"
      duplex                                   = "full"
      switchport_access_vlan                   = 100
      switchport_trunk_allowed_vlans           = "100,200"
      switchport_trunk_native_vlan             = 100
      dot1x_enable                             = true
      dot1x_port_control                       = "auto"
      dot1x_authentication_order               = ["dot1x"]
      dot1x_voice_vlan                         = 200
      dot1x_pae_enable                         = true
      dot1x_mac_authentication_bypass          = true
      dot1x_host_mode                          = "multi-domain"
      dot1x_enable_periodic_reauth             = true
      dot1x_periodic_reauth_inactivity_timeout = 100
      dot1x_periodic_reauth_interval           = 60
      dot1x_control_direction                  = "both"
      dot1x_restricted_vlan                    = 100
      dot1x_guest_vlan                         = 101
      dot1x_critical_vlan                      = 102
      dot1x_enable_criticial_voice_vlan        = true
    }
  ]
  age_out_time = 500
  static_mac_addresses = [
    {
      mac_address = "0000.0000.0000"
      if_name     = "GigabitEthernet0/0/0"
      vlan        = 100
    }
  ]
}
