resource "sdwan_cisco_snmp_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  shutdown     = false
  contact      = "Max"
  location     = "Building 1"
  views = [
    {
      name = "VIEW1"
      object_identifiers = [
        {
          id      = "1.2.3"
          exclude = true
        }
      ]
    }
  ]
  communities = [
    {
      name          = "community1"
      view          = "VIEW1"
      authorization = "read-only"
    }
  ]
  groups = [
    {
      name           = "GROUP1"
      security_level = "auth-priv"
      view           = "VIEW1"
    }
  ]
  users = [
    {
      name                    = "user1"
      authentication_protocol = "sha"
      authentication_password = "password123"
      privacy_protocol        = "aes-cfb-128"
      privacy_password        = "password123"
      group                   = "GROUP1"
    }
  ]
  trap_targets = [
    {
      vpn_id           = 1
      ip               = "1.1.1.1"
      udp_port         = 12345
      community_name   = "community1"
      user             = "user1"
      source_interface = "e1"
    }
  ]
}
