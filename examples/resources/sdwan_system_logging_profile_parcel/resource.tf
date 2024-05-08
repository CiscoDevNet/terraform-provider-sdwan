resource "sdwan_system_logging_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  disk_enable        = true
  disk_file_size     = 9
  disk_file_rotate   = 10
  tls_profiles = [
    {
      profile       = "test"
      tls_version   = "TLSv1.1"
      cipher_suites = ["aes-128-cbc-sha"]
    }
  ]
  ipv4_servers = [
    {
      hostname_ip                   = "1.1.1.1"
      vpn                           = 512
      source_interface              = "GigabitEthernet1"
      priority                      = "informational"
      tls_enable                    = true
      tls_properties_custom_profile = true
      tls_properties_profile        = "test"
    }
  ]
  ipv6_servers = [
    {
      hostname_ip                   = "1.1.1.1"
      vpn                           = 512
      source_interface              = "GigabitEthernet1"
      priority                      = "informational"
      tls_enable                    = true
      tls_properties_custom_profile = true
      tls_properties_profile        = "test"
    }
  ]
}
