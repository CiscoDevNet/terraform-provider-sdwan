resource "sdwan_policy_object_mirror" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  entries = [
    {
      remote_destination_ip = "10.0.0.1"
      source_ip             = "10.0.0.2"
    }
  ]
}
