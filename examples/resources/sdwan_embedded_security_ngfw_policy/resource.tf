resource "sdwan_embedded_security_ngfw_policy" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_actione    = "pass"
  rules = [
    {
      sequence_id  = "1"
      rule_name    = "Rule1"
      base_action  = "drop"
      rule_type    = "ngfirewall"
      disable_rule = false
      matches = [
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
