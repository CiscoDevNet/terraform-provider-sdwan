resource "sdwan_policy_object_expanded_community_profile_parcel" "example" {
  name                     = "Example"
  description              = "My Example"
  feature_profile_id       = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  expanded_community_lists = ["abcd"]
}
