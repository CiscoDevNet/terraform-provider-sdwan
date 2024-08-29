resource "sdwan_policy_object_as_path_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  as_path_list_id    = 1
  entries = [
    {
      as_path_list = "110"
    }
  ]
}
