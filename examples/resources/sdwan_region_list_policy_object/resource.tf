resource "sdwan_region_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      region_id = "1-2"
    }
  ]
}
