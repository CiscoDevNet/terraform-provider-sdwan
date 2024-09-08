resource "sdwan_service_wireless_lan_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  enable_24g         = true
  enable_5g          = true
  ssids = [
    {
      ssid_name      = "SSID_1"
      admin_state    = true
      broadcast_ssid = true
      vlan_id        = 1
      radio_type     = "all"
      security_type  = "personal"
      passphrase     = "MyPassword123"
      qos_profile    = "silver"
    }
  ]
  country               = "GB"
  username              = "user1"
  password              = "Test@316s13"
  me_dynamic_ip_enabled = true
}
