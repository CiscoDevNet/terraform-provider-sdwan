resource "sdwan_cellular_cedge_profile_feature_template" "example" {
  name                     = "Example"
  description              = "My Example"
  device_types             = ["vedge-C8000V"]
  profile_id               = 1
  access_point_name        = "APN1"
  authentication_type      = "chap"
  packet_data_network_type = "ipv4"
  profile_username         = "MyUsername"
  profile_password         = "MyPassword"
  no_overwrite             = false
}
