resource "sdwan_system_ntp_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  servers = [
    {
      hostname_ip_address    = "1.1.1.1"
      authentication_key     = 41673
      vpn                    = 1
      ntp_version            = 4
      source_interface       = "Ethernet"
      prefer_this_ntp_server = false
    }
  ]
  authentication_keys = [
    {
      key_id    = 49737
      md5_value = "$CRYPT_CLUSTER"
    }
  ]
  trusted_keys             = [49737]
  authoritative_ntp_server = false
  stratum                  = 1
  source_interface         = "ATM"
}
