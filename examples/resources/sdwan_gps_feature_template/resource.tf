resource "sdwan_gps_feature_template" "example" {
  name                = "Example"
  description         = "My Example"
  device_types        = ["vedge-C8000V"]
  enable              = true
  gps_mode            = "ms-based"
  nmea                = true
  source_address      = "1.2.3.4"
  destination_address = "2.3.4.5"
  destination_port    = 1234
}
