resource "sdwan_sse_zscaler_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  sse_provider       = "Zscaler"
  interfaces = [
    {
      interface_name          = "ipsec1"
      auto                    = true
      shutdown                = false
      tunnel_source_interface = "GigabitEthernet8"
      tunnel_set              = "secure-internet-gateway-zscaler"
      tunnel_dc_preference    = "primary-dc"
      mtu                     = 1400
      dpd_interval            = 10
      dpd_retries             = 3
      ike_version             = 2
      ike_rekey_interval      = 14400
      ike_ciphersuite         = "aes256-cbc-sha1"
      ike_group               = "16"
      ipsec_rekey_interval    = 3600
      ipsec_replay_window     = 512
      ipsec_ciphersuite       = "aes256-gcm"
      perfect_forward_secrecy = "none"
      track_enable            = true
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
  auth_required             = false
  xff_forward_enabled       = false
  ofw_enabled               = false
  ips_control               = false
  caution_enabled           = false
  refresh_time              = 1
  refresh_time_unit         = "MINUTE"
  aup_enabled               = false
  location_name             = "Auto"
  country                   = false
  enforce_bandwidth_control = true
  dn_bandwidth              = 0.1
  up_bandwidth              = 0.1
  sub_locations = [
    {
      name            = "zscaler_sub1"
      auth_required   = false
      ofw_enabled     = false
      caution_enabled = false
      service_vpn     = ["service_lan_vpn1,service_lan_vpn2"]
      internal_ip = [
        {
          internal_ip_value = "172.16.2.1"
        }
      ]
      aup_enabled               = false
      enforce_bandwidth_control = "location-bandwidth"
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
