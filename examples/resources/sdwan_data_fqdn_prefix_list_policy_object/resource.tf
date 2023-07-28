resource "sdwan_data_fqdn_prefix_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      fqdn = "cisco.com"
    }
  ]
}
