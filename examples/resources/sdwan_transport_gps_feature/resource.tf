resource "sdwan_transport_gps_feature" "example" {
  name                     = "Example"
  description              = "My Example"
  feature_profile_id       = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  gps_enable               = false
  gps_mode                 = "ms-based"
  nmea_enable              = false
  nmea_source_address      = "1.2.3.4"
  nmea_destination_address = "2.3.4.5"
  nmea_destination_port    = 22
}
