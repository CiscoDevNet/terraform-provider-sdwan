resource "sdwan_rewrite_rule_policy_definition" "example" {
  name        = "Example"
  description = "My description"
  rules = [
    {
      class_map_id = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
      priority     = "low"
      dscp         = 16
      layer2cos    = 1
    }
  ]
}
