resource "sdwan_ipv4_prefix_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      prefix = "10.0.0.0/12"
      le     = 32
      ge     = 24
    }
  ]
}
