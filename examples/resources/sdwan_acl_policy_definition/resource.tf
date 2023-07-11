resource "sdwan_acl_policy_definition" "example" {
  name           = "Example"
  description    = "My description"
  default_action = "drop"
  sequences = [
    {
      id          = 10
      ip_type     = "ipv4"
      name        = "Sequence 10"
      base_action = "accept"
      match_entries = [
        {
          type = "dscp"
          dscp = 16
        }
      ]
      action_entries = [
        {
          type     = "set"
          next_hop = "10.1.1.2"
        }
      ]
    }
  ]
}
