resource "sdwan_cli_config_profile_parcel" "example" {
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  name               = "Example"
  description        = "My Example"
  cli_configuration  = "bfd default-dscp 48\nbfd app-route multiplier 6\nbfd app-route poll-interval 600000"
}
