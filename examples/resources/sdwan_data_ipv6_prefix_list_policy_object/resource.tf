resource "sdwan_data_ipv6_prefix_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      prefix = "2001:0:0:1::/64"
    }
  ]
}
