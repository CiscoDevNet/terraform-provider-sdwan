resource "sdwan_rule_set_policy_definition" "example" {
  name        = "Example"
  description = "My description"
  rules = [
    {
      name                     = "Rule1"
      order                    = 1
      source_ipv4_prefix       = "10.1.1.0/24"
      source_fqdn              = "cisco.com"
      source_port              = "80-90"
      source_geo_location      = "AF"
      destination_ipv4_prefix  = "10.1.1.0/24"
      destination_fqdn         = "cisco.com"
      destination_port         = "80-90"
      destination_geo_location = "AF"
      protocol                 = "cifs"
    }
  ]
}
