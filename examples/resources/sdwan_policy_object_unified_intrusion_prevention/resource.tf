resource "sdwan_policy_object_unified_intrusion_prevention" "example" {
  name                  = "Example"
  description           = "My Example"
  feature_profile_id    = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  signature_set         = "balanced"
  inspection_mode       = "detection"
  ips_signature_list_id = "2ad58d78-59ee-46d3-86dd-7b6b7ca09f38"
  log_level             = "error"
  custom_signature      = false
}
