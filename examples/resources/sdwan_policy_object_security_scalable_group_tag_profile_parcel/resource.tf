resource "sdwan_policy_object_security_scalable_group_tag_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = ""
  entries = [
    {
      sgt_name = "ANY"
      tag      = "65535"
    }
  ]
}
