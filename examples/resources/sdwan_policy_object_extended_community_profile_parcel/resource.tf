resource "sdwan_policy_object_extended_community_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  entries = [
    {
      extended_community = "soo 10.0.0.1:30"
    }
  ]
}
