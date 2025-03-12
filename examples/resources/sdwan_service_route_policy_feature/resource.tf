resource "sdwan_service_route_policy_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_action     = "accept"
  sequences = [
    {
      id          = 1
      name        = "SEQ_1"
      base_action = "reject"
      protocol    = "IPV4"
      actions = [
        {
          as_path_prepend    = [65521]
          community_additive = false
          community          = ["internet"]
          local_preference   = 100
          metric             = 20
          metric_type        = "type1"
          omp_tag            = 200
          origin             = "EGP"
          ospf_tag           = 1200
          weight             = 2200
          ipv4_next_hop      = "10.0.0.1"
        }
      ]
    }
  ]
}
