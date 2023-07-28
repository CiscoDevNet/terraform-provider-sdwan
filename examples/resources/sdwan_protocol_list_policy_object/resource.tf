resource "sdwan_protocol_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      protocol = "cifs"
    }
  ]
}
