resource "sdwan_hub_and_spoke_topology_policy_definition" "example" {
  name        = "Example"
  description = "My description"
  vpn_list_id = "04fcbb0b-efbf-43d2-a04b-847d3a7b104e"
  topologies = [
    {
      name                = "Topology1"
      all_hubs_are_equal  = false
      advertise_hub_tlocs = true
      tloc_list_id        = "b326e448-bf33-47e4-83e7-f947e6981382"
      spokes = [
        {
          site_list_id = "e858e1c4-6aa8-4de7-99df-c3adbf80290d"
          hubs = [
            {
              site_list_id = "e858e1c4-6aa8-4de7-99df-c3adbf80290d"
              preference   = "30"
            }
          ]
        }
      ]
    }
  ]
}
