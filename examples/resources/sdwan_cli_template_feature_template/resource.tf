resource "sdwan_cli_template_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  cli_config   = "! Enable new BGP community format\nip bgp-community new-format\n"
}
