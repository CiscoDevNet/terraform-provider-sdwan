resource "sdwan_service_tracker_group_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  tracker_elements = [
    {
      tracker_id = "615d948f-34ee-4a2e-810e-a9bd8d3d48ec"
    }
  ]
  tracker_boolean = "or"
}
