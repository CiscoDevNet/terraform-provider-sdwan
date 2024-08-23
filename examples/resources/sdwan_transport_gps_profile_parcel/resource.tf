resource "sdwan_transport_gps_profile_parcel" "example" {
  name                = "Example"
  description         = "My Example"
  feature_profile_id  = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  enable              = false
  mode                = "ms-based"
  nmea                = false
  source_address      = "1.2.3.4"
  destination_address = "2.3.4.5"
  destination_port    = 22
}
