resource "sdwan_object_group_policy_definition" "example" {
  name         = "Example"
  description  = "My description"
  ipv4_prefix  = "10.1.1.0/24"
  fqdn         = "cisco.com"
  port         = "80-90"
  geo_location = "AF"
}
