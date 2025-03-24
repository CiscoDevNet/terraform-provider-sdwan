resource "sdwan_dns_security_dns_security_policy" "example" {
  name                        = "Example"
  description                 = "My Example"
  feature_profile_id          = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  local_domain_bypass_list_id = "0c4e096d-b06d-4052-93ed-70fe34fda6dc"
  match_all_vpn               = true
  umbrella_default            = false
  dns_server_i_p              = "1.2.3.4"
  local_domain_bypass_enabled = true
  dns_crypt                   = false
  child_org_id                = "12334"
}
