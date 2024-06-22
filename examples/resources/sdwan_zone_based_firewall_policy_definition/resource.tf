resource "sdwan_zone_based_firewall_policy_definition" "example" {
  name        = "Example"
  description = "My description"
  mode        = "security"
  apply_zone_pairs = [
    {
      source_zone      = "self"
      destination_zone = "0d26a366-4a11-4942-a5ea-82af9502889f"
    }
  ]
  default_action = "pass"
  rules = [
    {
      rule_order  = 1
      rule_name   = "RULE_1"
      base_action = "inspect"
      match_entries = [
        {
          type      = "sourceGeoLocationList"
          policy_id = "0d26a366-4a11-4942-a5ea-82af9502889f"
        }
      ]
      action_entries = [
        {
          type = "log"
        }
      ]
    }
  ]
}
