resource "sdwan_policy_object_tloc_list_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  entries = [
    {
      tloc_ip       = "10.0.0.0"
      color         = "3g"
      encapsulation = "gre"
      preference    = "33"
    }
  ]
}
