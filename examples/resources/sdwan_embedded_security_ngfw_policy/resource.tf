resource "sdwan_embedded_security_ngfw_policy" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_action     = "pass"
  sequences = [
    {
      sequence_id   = "1"
      sequence_name = "security"
      base_action   = "drop"
      rule_type     = "ngfirewall"
      disable_rule  = false
      match_entries = [
        {
          source_ports = ["123"]
        }
      ]
      actions = [
        {
          type      = "log"
          parameter = "true"
        }
      ]
    }
  ]
}
