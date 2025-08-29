resource "sdwan_port_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      port = "80"
    }
  ]
}
