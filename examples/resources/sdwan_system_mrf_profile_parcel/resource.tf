resource "sdwan_system_mrf_profile_parcel" "example" {
  name                    = "Example"
  description             = "My Example"
  feature_profile_id      = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  region_id               = 1
  secondary_region        = 2
  role                    = "edge-router"
  enable_mrf_migration    = "enabled"
  migration_bgp_community = 100
}
