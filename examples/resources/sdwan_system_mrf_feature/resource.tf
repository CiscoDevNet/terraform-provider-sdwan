resource "sdwan_system_mrf_feature" "example" {
  name                     = "Example"
  description              = "My Example"
  feature_profile_id       = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  secondary_region_id      = 2
  role                     = "edge-router"
  enable_migration_to_mrf  = "enabled"
  migration_bgp_community  = 100
  enable_management_region = false
  vrf_id                   = false
  gateway_preference       = [1]
  management_gateway       = false
}
