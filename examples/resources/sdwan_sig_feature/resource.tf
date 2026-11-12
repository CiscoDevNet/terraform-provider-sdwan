resource "sdwan_sig_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  sig_provider       = "Generic"
  interfaces = [
    {
      interface_name          = "gre1"
      interface_description   = "Tunnel to SIG data center A"
      tunnel_source_interface = "GigabitEthernet1"
      tunnel_destination      = "198.51.100.10"
      tunnel_set              = "secure-internet-gateway-other"
      tunnel_dc_preference    = "primary-dc"
      tcp_mss_adjust          = 1300
      mtu                     = 1400
      tracker                 = "tracker1"
      track_enable            = true
    }
  ]
  interface_pairs = [
    {
      active_interface        = "gre1"
      active_interface_weight = 255
      backup_interface        = "None"
      backup_interface_weight = 1
    }
  ]
  tracker_source_ip = "192.0.2.10"
  trackers = [
    {
      name             = "tracker1"
      endpoint_api_url = "http://gateway.example.com/vpntest"
      threshold        = 1000
      interval         = 30
      multiplier       = 2
    }
  ]
}
