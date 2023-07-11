resource "sdwan_cisco_ntp_feature_template" "example" {
  name                    = "Example"
  description             = "My Example"
  device_types            = ["vedge-C8000V"]
  master                  = true
  master_stratum          = 6
  master_source_interface = "e1"
  trusted_keys            = [1]
  authentication_keys = [
    {
      id    = 1
      value = "12345"
    }
  ]
  servers = [
    {
      hostname_ip           = "NTP_SERVER1"
      authentication_key_id = 1
      vpn_id                = 1
      version               = 4
      source_interface      = "e1"
      prefer                = true
    }
  ]
}
