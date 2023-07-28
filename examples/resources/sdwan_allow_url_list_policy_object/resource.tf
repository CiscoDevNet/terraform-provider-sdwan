resource "sdwan_allow_url_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      url = "cisco.com"
    }
  ]
}
