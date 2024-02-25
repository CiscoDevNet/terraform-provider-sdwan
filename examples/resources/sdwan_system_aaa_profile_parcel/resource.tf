resource "sdwan_system_aaa_profile_parcel" "example" {
  name                 = "Example"
  description          = "My Example"
  feature_profile_id   = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  authentication_group = true
  accounting_group     = true
  server_auth_order    = ["local"]
  users = [
    {
      name      = "User1"
      password  = "cisco123"
      privilege = "15"
      public_keys = [
        {
          key_string = "AAAAB3NzaC1yc2"
          key_type   = "ssh-rsa"
        }
      ]
    }
  ]
  radius_groups = [
    {
      group_name       = "RGROUP1"
      vpn              = 10
      source_interface = "GigabitEthernet0"
      servers = [
        {
          address    = "1.2.3.4"
          auth_port  = 1812
          acct_port  = 1813
          timeout    = 5
          retransmit = 3
          key        = "cisco123"
          secret_key = "cisco123"
          key_enum   = "7"
          key_type   = "key"
        }
      ]
    }
  ]
  tacacs_groups = [
    {
      group_name       = "TGROUP1"
      vpn              = 10
      source_interface = "GigabitEthernet0"
      servers = [
        {
          address    = "1.2.3.4"
          port       = 49
          timeout    = 5
          key        = "cisco123"
          secret_key = "cisco123"
          key_enum   = "7"
        }
      ]
    }
  ]
  accounting_rules = [
    {
      rule_id    = "1"
      method     = "commands"
      level      = "15"
      start_stop = true
      group      = ["RGROUP1"]
    }
  ]
  authorization_console         = true
  authorization_config_commands = true
  authorization_rules = [
    {
      rule_id          = "1"
      method           = "commands"
      level            = "15"
      group            = ["RGROUP1"]
      if_authenticated = true
    }
  ]
}
