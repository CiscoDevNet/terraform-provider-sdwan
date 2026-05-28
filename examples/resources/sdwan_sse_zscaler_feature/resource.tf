resource "sdwan_sse_zscaler_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  interfaces = [
    {
      interface_name          = "ipsec1"
      auto                    = true
      tunnel_source_interface = "GigabitEthernet8"
      tunnel_set              = "secure-internet-gateway-zscaler"
      tunnel_dc_preference    = "primary-dc"
      mtu                     = 1400
      ike_version             = 2
      pre_shared_key_dynamic  = true
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
  refresh_time      = 1
  refresh_time_unit = "MINUTE"
  sub_locations = [
    {
      name            = "zscaler_sub1"
      ofw_enabled     = false
      caution_enabled = false
      service_vpn     = ["service_lan_vpn1,service_lan_vpn2"]
      internal_ip = [
        {
          internal_ip_value = "172.16.2.1"
        }
      ]
      aup_enabled               = false
      enforce_bandwidth_control = "disabled"
    }
  ]
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
