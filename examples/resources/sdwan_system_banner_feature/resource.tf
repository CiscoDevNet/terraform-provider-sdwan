resource "sdwan_system_banner_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  login              = "My login banner"
  motd               = "My motd banner"
}
