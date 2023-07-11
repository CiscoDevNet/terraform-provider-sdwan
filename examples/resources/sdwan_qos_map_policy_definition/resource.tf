resource "sdwan_qos_map_policy_definition" "example" {
  name        = "Example"
  description = "My description"
  qos_schedulers = [
    {
      queue             = 6
      class_map_id      = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
      bandwidth_percent = 10
      buffer_percent    = 10
      burst             = 100000
      drop_type         = "red-drop"
      scheduling_type   = "wrr"
    }
  ]
}
