resource "sdwan_topology_hub_spoke_feature" "example" {
  name                    = "Example"
  description             = "My Example"
  feature_profile_id      = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  target_vpns             = ["service_lan_vpn1"]
  selected_hierarchy_hubs = ["acb2ea53-4a95-4970-a1ab-9bac15edb961"]
  spokes = [
    {
      name                  = "spoke1"
      spoke_hierarchy_uuids = ["acb2ea53-4a95-4970-a1ab-9bac15edb961"]
      hub_sites = [
        {
          hub_hierarchy_uuids = ["acb2ea53-4a95-4970-a1ab-9bac15edb961"]
          preference          = 1
        }
      ]
    }
  ]
}
