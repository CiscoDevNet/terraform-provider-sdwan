resource "sdwan_service_dual_router_ha_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  redundancy_groups = [
    {
      group_id = 1
      vpn_ids = [
        {
          vpn_id = "615d948f-34ee-4a2e-810e-a9bd8d3d48ec"
        }
      ]
      tag_name = "example"
    }
  ]
  enable_optimize_paths = false
}
