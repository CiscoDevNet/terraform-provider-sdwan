resource "sdwan_cisco_logging_feature_template" "example" {
  name          = "Example"
  description   = "My Example"
  device_types  = ["vedge-C8000V"]
  disk_logging  = true
  max_size      = 10
  log_rotations = 10
  tls_profiles = [
    {
      name                = "PROF1"
      version             = "TLSv1.2"
      authentication_type = "Server"
      ciphersuite_list    = ["aes-128-cbc-sha"]
    }
  ]
  ipv4_servers = [
    {
      hostname_ip      = "2.2.2.2"
      vpn_id           = 1
      source_interface = "e1"
      logging_level    = "information"
      enable_tls       = true
      custom_profile   = true
      profile          = "PROF1"
    }
  ]
  ipv6_servers = [
    {
      hostname_ip      = "2001::1"
      vpn_id           = 1
      source_interface = "e1"
      logging_level    = "information"
      enable_tls       = true
      custom_profile   = true
      profile          = "PROF1"
    }
  ]
}
