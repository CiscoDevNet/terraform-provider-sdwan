resource "sdwan_feature_device_template" "example" {
  name        = "Example"
  description = "My description"
  device_type = "vedge-ISR-4331"
  general_templates = [
    {
      id   = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
      type = "cisco_system"
    }
  ]
}
