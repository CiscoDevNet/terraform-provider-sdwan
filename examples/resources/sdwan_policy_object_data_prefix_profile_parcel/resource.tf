resource "sdwan_policy_object_data_prefix_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  entries = [
    {
      ipv4_address       = "10.0.0.0"
      ipv4_prefix_length = 8
    }
  ]
}
