resource "sdwan_route_policy_definition" "example" {
  name           = "Example"
  description    = "My description"
  default_action = "reject"
  sequences = [
    {
      id          = 10
      ip_type     = "ipv4"
      name        = "Sequence 10"
      base_action = "accept"
      match_entries = [
        {
          type   = "metric"
          metric = 100
        }
      ]
      action_entries = [
        {
          type                  = "aggregator"
          aggregator            = 10
          aggregator_ip_address = "10.1.2.3"
        }
      ]
    }
  ]
}
