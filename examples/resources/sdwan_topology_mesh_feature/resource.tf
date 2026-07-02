resource "sdwan_topology_mesh_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  target_vpns        = ["service_lan_vpn1"]
  sites              = ["SITE_100"]
}
