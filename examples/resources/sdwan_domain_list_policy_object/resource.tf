resource "sdwan_domain_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      domain = ".*.cisco.com"
    }
  ]
}
