resource "sdwan_transport_cellular_profile_feature" "example" {
  name                     = "Example"
  description              = "My Example"
  feature_profile_id       = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  profile_id               = 1
  access_point_name        = "apn1"
  need_authentication      = "pap"
  packet_data_network_type = "ipv4"
  no_overwrite             = false
}
