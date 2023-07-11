resource "sdwan_cedge_aaa_feature_template" "example" {
  name                         = "Example"
  description                  = "My Example"
  device_types                 = ["vedge-C8000V"]
  dot1x_authentication         = true
  dot1x_accounting             = true
  server_groups_priority_order = "100"
  users = [
    {
      name            = "user1"
      password        = "password123"
      secret          = "secret123"
      privilege_level = "15"
      ssh_pubkeys = [
        {
          key_string = "abc123"
          key_type   = "rsa"
        }
      ]
    }
  ]
  radius_server_groups = [
    {
      group_name       = "GROUP1"
      vpn_id           = 1
      source_interface = "e1"
      servers = [
        {
          address             = "1.1.1.1"
          authentication_port = 1812
          accounting_port     = 1813
          timeout             = 5
          retransmit          = 3
          key                 = "key123"
          secret_key          = "1234567"
          encryption_type     = "7"
          key_type            = "pac"
        }
      ]
    }
  ]
  radius_clients = [
    {
      client_ip = "2.2.2.2"
      von_configurations = [
        {
          vpn_id     = "1"
          server_key = "key123"
        }
      ]
    }
  ]
  radius_dynamic_author_server_key             = "key123"
  radius_dynamic_author_domain_stripping       = "yes"
  radius_dynamic_author_authentication_type    = "all"
  radius_dynamic_author_port                   = 1700
  radius_dynamic_author_cts_authorization_list = "ALIST1"
  radius_trustsec_group                        = "GROUP1"
  tacacs_server_groups = [
    {
      group_name       = "GROUP1"
      vpn_id           = 1
      source_interface = "e1"
      servers = [
        {
          address         = "1.1.1.1"
          port            = 49
          timeout         = 5
          key             = "key123"
          secret_key      = "1234567"
          encryption_type = "7"
        }
      ]
    }
  ]
  accounting_rules = [
    {
      name            = "RULE1"
      method          = "exec"
      privilege_level = "15"
      start_stop      = true
      group           = ["GROUP1"]
    }
  ]
  authorization_console         = true
  authorization_config_commands = true
  authorization_rules = [
    {
      name            = "RULE1"
      method          = "commands"
      privilege_level = "15"
      group           = ["GROUP1"]
      authenticated   = true
    }
  ]
}
