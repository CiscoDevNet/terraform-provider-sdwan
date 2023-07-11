resource "sdwan_cisco_security_feature_template" "example" {
  name                = "Example"
  description         = "My Example"
  device_types        = ["vedge-C8000V"]
  rekey_interval      = 86400
  replay_window       = "64"
  extended_ar_window  = 256
  authentication_type = ["none"]
  integrity_type      = ["none"]
  pairwise_keying     = true
  keychains = [
    {
      name   = "CHAIN1"
      key_id = 1
    }
  ]
  keys = [
    {
      id                              = "1"
      chain_name                      = "CHAIN1"
      send_id                         = 0
      receive_id                      = 0
      crypto_algorithm                = "hmac-sha-256"
      key_string                      = "abc123"
      send_lifetime                   = true
      send_lifetime_start_time        = "2022-12-31T23:59"
      send_lifetime_end_time_format   = "infinite"
      send_lifetime_duration          = 1000
      send_lifetime_end_time          = "2032-12-31T23:59"
      send_lifetime_infinite          = true
      accept_lifetime                 = true
      accept_lifetime_start_time      = "2022-12-31T23:59"
      accept_lifetime_end_time_format = "infinite"
      accept_lifetime_duration        = 1000
      accept_lifetime_end_time        = "2032-12-31T23:59"
      accept_lifetime_infinite        = true
      include_tcp_options             = false
      accept_ao_mismatch              = true
    }
  ]
}
