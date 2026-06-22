resource "sdwan_topology_hub_spoke_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  target_vpns        = ["service_lan_vpn1"]
  selected_hubs      = ["SITE_100"]
  spokes = [
    {
      name        = "spoke1"
      spoke_sites = ["SITE_200"]
      hub_sites = [
        {
          sites      = ["SITE_100"]
          preference = 1
        }
      ]
    }
  ]
}
