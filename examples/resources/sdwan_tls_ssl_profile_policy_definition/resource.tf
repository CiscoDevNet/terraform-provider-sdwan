resource "sdwan_tls_ssl_profile_policy_definition" "example" {
  name                     = "Example"
  description              = "My description"
  mode                     = "security"
  decrypt_categories       = ["alcohol-and-tobacco"]
  never_decrypt_categories = ["auctions"]
  skip_decrypt_categories  = ["cdns"]
  decrypt_threshold        = "high-risk"
  reputation               = false
  fail_decrypt             = true
}
