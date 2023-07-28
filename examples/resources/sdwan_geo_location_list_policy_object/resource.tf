resource "sdwan_geo_location_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      country   = "USA"
      continent = "AS"
    }
  ]
}
