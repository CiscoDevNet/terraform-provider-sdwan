resource "sdwan_data_ipv4_prefix_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      prefix = "10.0.0.0/12"
    }
  ]
}
