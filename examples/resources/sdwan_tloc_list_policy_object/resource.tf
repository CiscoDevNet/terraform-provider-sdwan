resource "sdwan_tloc_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      tloc_ip       = "1.1.1.2"
      color         = "blue"
      encapsulation = "gre"
      preference    = 10
    }
  ]
}
