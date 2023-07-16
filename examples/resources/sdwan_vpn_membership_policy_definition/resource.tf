resource "sdwan_vpn_membership_policy_definition" "example" {
  name        = "Example"
  description = "My description"
  sites = [
    {
      site_list_id = "e858e1c4-6aa8-4de7-99df-c3adbf80290d"
      vpn_list_ids = ["04fcbb0b-efbf-43d2-a04b-847d3a7b104e"]
    }
  ]
}
