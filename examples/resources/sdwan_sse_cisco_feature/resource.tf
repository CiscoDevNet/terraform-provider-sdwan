resource "sdwan_sse_cisco_feature" "example" {
  name                    = "Example"
  description             = "My Example"
  feature_profile_id      = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  context_sharing_for_vpn = false
  context_sharing_for_sgt = false
  interfaces = [
    {
      interface_name          = "ipsec1"
      tunnel_source_interface = "GigabitEthernet8"
      tunnel_dc_preference    = "primary-dc"
      mtu                     = 1400
      ike_version             = 2
    }
  ]
  interface_pairs = [
    {
      active_interface        = "ipsec1"
      active_interface_weight = 1
      backup_interface        = "None"
      backup_interface_weight = 1
    }
  ]
  region            = "us-east-1"
  tracker_source_ip = "1.2.3.4"
  trackers = [
    {
      name             = "tracker1"
      endpoint_api_url = "http://cisco.com"
      threshold        = 300
      interval         = 60
      multiplier       = 3
    }
  ]
}
