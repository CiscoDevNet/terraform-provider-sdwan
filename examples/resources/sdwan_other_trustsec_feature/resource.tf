resource "sdwan_other_trustsec_feature" "example" {
  name                      = "Example"
  description               = "My Example"
  feature_profile_id        = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  device_id                 = "trustsecDevice"
  device_password           = "Cisco123"
  device_sgt                = 100
  enable_enforcement        = true
  enable_sxp                = true
  listener_hold_time_min    = 90
  listener_hold_time_max    = 180
  speaker_hold_time         = 120
  sxp_key_chain             = "keychain1"
  sxp_log_binding_changes   = true
  sxp_reconciliation_period = 120
  sxp_retry_period          = 120
  sxp_source_ip             = "10.0.0.1"
  sxp_connections = [
    {
      peer_ip       = "10.0.0.2"
      source_ip     = "10.0.0.1"
      preshared_key = "none"
      mode          = "peer"
      mode_type     = "speaker"
      vpn_id        = 1
      min_hold_time = 90
      max_hold_time = 120
    }
  ]
}
