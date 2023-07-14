resource "sdwan_cisco_secure_internet_gateway_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  vpn_id       = 1
  interfaces = [
    {
      name                          = "ipsec1"
      auto_tunnel_mode              = true
      shutdown                      = true
      description                   = "My Description"
      ip_unnumbered                 = true
      ipv4_address                  = "1.2.3.4/24"
      tunnel_source                 = "3.3.3.3"
      tunnel_source_interface       = "ge0/1"
      tunnel_route_via              = "ge0/2"
      tunnel_destination            = "3.4.5.6"
      application                   = "sig"
      sig_provider                  = "secure-internet-gateway-umbrella"
      tunnel_dc_preference          = "primary-dc"
      tcp_mss                       = 1400
      mtu                           = 1500
      dead_peer_detection_interval  = 30
      dead_peer_detection_retries   = 5
      ike_version                   = 1
      ike_pre_shared_key            = "A1234567"
      ike_rekey_interval            = 600
      ike_ciphersuite               = "aes256-cbc-sha2"
      ike_group                     = "14"
      ike_pre_shared_key_dynamic    = false
      ike_pre_shared_key_local_id   = "1.2.3.4"
      ike_pre_shared_key_remote_id  = "2.3.4.5"
      ipsec_rekey_interval          = 7200
      ipsec_replay_window           = 1024
      ipsec_ciphersuite             = "aes256-cbc-sha1"
      ipsec_perfect_forward_secrecy = "group-14"
      track_enable                  = false
      tunnel_public_ip              = "5.5.5.5"
    }
  ]
  services = [
    {
      service_type = "sig"
      interface_pairs = [
        {
          active_interface        = "e1"
          active_interface_weight = 10
          backup_interface        = "e2"
          backup_interface_weight = 20
        }
      ]
      zscaler_authentication_required                 = true
      zscaler_xff_forward                             = true
      zscaler_firewall_enabled                        = true
      zscaler_ips_control_enabled                     = true
      zscaler_caution_enabled                         = true
      zscaler_primary_data_center                     = "Auto"
      zscaler_secondary_data_center                   = "Auto"
      zscaler_surrogate_ip                            = true
      zscaler_surrogate_idle_time                     = 100
      zscaler_surrogate_display_time_unit             = "MINUTE"
      zscaler_surrogate_ip_enforce_for_known_browsers = true
      zscaler_surrogate_refresh_time_unit             = "MINUTE"
      aup_enabled                                     = true
      aup_block_internet_until_accepted               = true
      aup_force_ssl_inspection                        = true
      aup_timeout                                     = 60
      zscaler_location_name                           = "LOC1"
      umbrella_primary_data_center                    = "Auto"
      umbrella_secondary_data_center                  = "Auto"
    }
  ]
  tracker_source_ip = "2.3.4.5"
  trackers = [
    {
      name             = "TRACKER1"
      endpoint_api_url = "https://1.1.1.1"
      threshold        = 500
      multiplier       = 4
      tracker_type     = "SIG"
    }
  ]
}
