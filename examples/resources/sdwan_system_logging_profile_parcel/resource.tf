resource "sdwan_system_logging_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  disk_enable        = true
  disk_file_size     = 9
  disk_file_rotate   = 10
  tls_profile = [
    {
      profile           = "test"
      tls_version       = "TLSv1.1"
      auth_type         = "Server"
      cipher_suite_list = ["aes-128-cbc-sha"]
    }
  ]
  server = [
    {
      name                          = "1.1.1.1"
      vpn                           = 512
      source_interface              = "GigabitEthernet1"
      priority                      = "informational"
      tls_enable                    = true
      tls_properties_custom_profile = false
      tls_properties_profile        = "test"
    }
  ]
  ipv6_server = [
    {
      name                          = "1.1.1.1"
      vpn                           = 512
      source_interface              = "GigabitEthernet1"
      priority                      = "informational"
      tls_enable                    = true
      tls_properties_custom_profile = false
      tls_properties_profile        = "test"
    }
  ]
}
