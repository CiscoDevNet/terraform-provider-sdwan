resource "sdwan_zone_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      vpn = "1"
    }
  ]
}
