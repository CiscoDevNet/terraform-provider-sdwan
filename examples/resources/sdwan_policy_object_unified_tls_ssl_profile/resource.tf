resource "sdwan_policy_object_unified_tls_ssl_profile" "example" {
  name                    = "Example"
  description             = "My Example"
  feature_profile_id      = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  decrypt_categories      = ["alcohol-and-tobacco"]
  no_decrypt_categories   = ["abortion"]
  pass_through_categories = ["auctions"]
  reputation              = true
  decrypt_threshold       = "moderate-risk"
  threshold_categories    = "moderate-risk"
  fail_decrypt            = true
  url_allow_list_id       = "2ad58d78-59ee-46d3-86dd-7b6b7ca09f38"
  url_block_list_id       = "2ad58d78-59ee-46d3-86dd-7b6b7ca09f38"
}
