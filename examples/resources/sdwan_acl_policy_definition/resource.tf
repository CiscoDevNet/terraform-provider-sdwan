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
          type = "set"
          set_parameters = [
            {
              type = "dscp"
              dscp = 16
            }
          ]
        }
      ]
    }
  ]
}
