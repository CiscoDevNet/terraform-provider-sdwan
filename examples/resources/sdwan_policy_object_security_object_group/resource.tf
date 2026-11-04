resource "sdwan_policy_object_security_object_group" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  sequence_ip_type   = "ipv4"
  entries = [
    {
      data_ipv4_prefixes = ["10.1.1.0/24"]
      fqdns              = ["cisco.com"]
      geo_locations      = ["EU"]
    }
  ]
}
