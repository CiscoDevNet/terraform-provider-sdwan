resource "sdwan_mesh_topology_policy_definition" "example" {
  name        = "Example"
  description = "My description"
  vpn_list_id = "04fcbb0b-efbf-43d2-a04b-847d3a7b104e"
  regions = [
    {
      name          = "Region1"
      site_list_ids = ["e858e1c4-6aa8-4de7-99df-c3adbf80290d"]
    }
  ]
}
