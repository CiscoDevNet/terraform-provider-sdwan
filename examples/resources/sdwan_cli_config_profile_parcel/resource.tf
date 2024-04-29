resource "sdwan_cli_config_profile_parcel" "example" {
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  name               = "Example"
  description        = "My Example"
  cli_config         = "! Enable new BGP community format\nip bgp-community new-format\n"
}
