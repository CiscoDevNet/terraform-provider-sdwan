resource "sdwan_custom_control_topology_policy_definition" "example" {
  name           = "Example"
  description    = "My description"
  default_action = "reject"
  sequences = [
    {
      id          = 1
      name        = "Region1"
      type        = "route"
      ip_type     = "ipv4"
      base_action = "accept"
      match_entries = [
        {
          type    = "ompTag"
          omp_tag = 100
        }
      ]
      action_entries = [
        {
          type = "set"
          set_parameters = [
            {
              type       = "preference"
              preference = 100
            }
          ]
        }
      ]
    }
  ]
}
