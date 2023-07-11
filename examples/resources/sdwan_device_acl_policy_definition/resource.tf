resource "sdwan_device_acl_policy_definition" "example" {
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
          type             = "destinationPort"
          destination_port = 22
        }
      ]
      action_entries = [
        {
          type         = "count"
          counter_name = "count1"
        }
      ]
    }
  ]
}
