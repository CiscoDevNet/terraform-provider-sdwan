resource "sdwan_policy_object_preferred_color_group_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  entries = [
    {
      primary_color_preference   = ["default"]
      primary_path_preference    = "direct-path"
      secondary_color_preference = ["bronze"]
      secondary_path_preference  = "all-paths"
      tertiary_color_preference  = ["blue"]
      tertiary_path_preference   = "all-paths"
    }
  ]
}
