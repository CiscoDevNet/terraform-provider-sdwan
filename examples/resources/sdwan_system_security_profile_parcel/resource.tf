resource "sdwan_system_security_profile_parcel" "example" {
  name                        = "Example"
  description                 = "My Example"
  feature_profile_id          = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  rekey                       = 86400
  anti_replay_window          = "512"
  extended_anti_replay_window = 256
  integrity_type              = ["esp"]
  ipsec_pairwise_keying       = false
  keychains = [
    {
      key_chain_name = "aaa"
      key_id         = 1
    }
  ]
  key = [
    {
      id                           = 0
      name                         = "aaa"
      send_id                      = 1
      receiver_id                  = 2
      include_tcp_options          = false
      accept_ao_mismatch           = false
      crypto_algorithm             = "aes-128-cmac"
      key_string                   = "abcabc"
      send_life_time_local         = true
      send_life_time_start_epoch   = 1659284400
      send_life_time_infinite      = true
      send_life_time_duration      = 2147483645
      send_life_time_exact         = 1
      accept_life_time_local       = true
      accept_life_time_start_epoch = 1659284400
      accept_life_time_infinite    = true
      accept_life_time_duration    = 2147483645
      accept_life_time_exact       = 1
    }
  ]
}
