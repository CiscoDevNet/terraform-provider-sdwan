resource "sdwan_policy_object_unified_advanced_inspection_profile" "example" {
  name                                = "Example"
  description                         = "My Example"
  feature_profile_id                  = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  tls_decryption_action               = "decrypt"
  intrusion_prevention_list_id        = "2ad58d78-59ee-46d3-86dd-7b6b7ca09f38"
  url_filtering_list_id               = "2ad58d78-59ee-46d3-86dd-7b6b7ca09f38"
  advanced_malware_protection_list_id = "2ad58d78-59ee-46d3-86dd-7b6b7ca09f38"
  tls_ssl_profile_list_id             = "2ad58d78-59ee-46d3-86dd-7b6b7ca09f38"
}
