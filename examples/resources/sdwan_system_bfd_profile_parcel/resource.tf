resource "sdwan_system_bfd_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  multiplier         = 3
  poll_interval      = 100
  default_dscp       = 8
  colors = [
    {
      color          = "3g"
      hello_interval = 200
      multiplier     = 3
      pmtu_discovery = true
      dscp           = 16
    }
  ]
}
