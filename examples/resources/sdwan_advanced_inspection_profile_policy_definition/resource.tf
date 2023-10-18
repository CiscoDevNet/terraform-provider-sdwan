resource "sdwan_advanced_inspection_profile_policy_definition" "example" {
  name             = "Example"
  description      = "My description"
  tls_action       = "decrypt"
  url_filtering_id = "914670a3-9726-4a51-847f-b3db70819dc2"
}
