resource "sdwan_cisco_banner_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  login        = "My Login Banner"
  motd         = "My MOTD Banner"
}
