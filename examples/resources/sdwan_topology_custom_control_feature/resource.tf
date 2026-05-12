resource "sdwan_topology_custom_control_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_action     = "reject"
  sequences = [
    {
      id          = 1
      name        = "Rule1"
      base_action = "accept"
      type        = "route"
      ip_type     = "ipv4"
      match_entries = [
        {
          omp_tag = 100
        }
      ]
      action_entries = [
        {
          set_parameters = [
            {
              preference = 100
            }
          ]
        }
      ]
    }
  ]
}
