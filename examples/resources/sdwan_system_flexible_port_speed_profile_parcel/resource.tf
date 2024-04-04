resource "sdwan_system_flexible_port_speed_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  port_type          = "12 ports of 1/10GE + 3 ports 40GE"
}
