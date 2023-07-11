resource "sdwan_ipv6_prefix_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      prefix = "2001:1:1:2::/64"
      le     = 80
      ge     = 128
    }
  ]
}
