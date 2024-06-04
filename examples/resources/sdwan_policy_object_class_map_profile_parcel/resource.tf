resource "sdwan_policy_object_class_map_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "e4d9392a-7765-4a64-b719-a4bcaf534f25"
  entries = [
    {
      queue = "0"
    }
  ]
}
