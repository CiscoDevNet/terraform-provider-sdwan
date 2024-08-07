resource "sdwan_policy_object_data_ipv6_prefix_list_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  entries = [
    {
      ipv6_address       = "2001:db8:85a3::8a2e:370:7334"
      ipv6_prefix_length = 64
    }
  ]
}
