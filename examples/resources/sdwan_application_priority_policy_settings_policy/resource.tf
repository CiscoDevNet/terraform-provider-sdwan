resource "sdwan_application_priority_policy_settings_policy" "example" {
  name                        = "Example"
  description                 = "My Example"
  feature_profile_id          = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  ipv4_application_visibility = true
  ipv4_flow_visibility        = true
}
