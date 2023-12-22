resource "sdwan_ipv6_acl_policy_definition" "example" {
  name           = "Example"
  description    = "My description"
  default_action = "drop"
  sequences = [
    {
      id          = 10
      name        = "Sequence 10"
      base_action = "accept"
      match_entries = [
        {
          type          = "nextHeader"
          next_header   = 1
          traffic_class = 1
        }
      ]
      action_entries = [
        {
          type = "set"
          set_parameters = [
            {
              type          = "trafficClass"
              traffic_class = 16
            }
          ]
        }
      ]
    }
  ]
}
