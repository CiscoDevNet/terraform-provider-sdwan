resource "sdwan_site_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      site_id = "100-200"
    }
  ]
}
