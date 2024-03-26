resource "sdwan_cisco_wireless_lan_feature_template" "example" {
  name            = "Example"
  description     = "My Example"
  device_types    = ["vedge-C8000V"]
  shutdown_2_4ghz = false
  shutdown_5ghz   = false
  ssids = [
    {
      wireless_network_name = "WLAN1"
      admin_state           = false
      broadcast_ssid        = true
      vlan_id               = 1
      radio_type            = "24ghz"
      security_type         = "enterprise"
      radius_server_ip      = "1.2.3.4"
      radius_server_port    = 1812
      radius_server_secret  = "MySecret1"
      passphrase            = "passphrase"
      qos_profile           = "silver"
    }
  ]
  country                    = "AE"
  username                   = "user1"
  password                   = "myPassword01"
  controller_ip_address      = "0.0.0.0"
  controller_subnet_mask     = "0.0.0.0"
  controller_default_gateway = "0.0.0.0"
}
