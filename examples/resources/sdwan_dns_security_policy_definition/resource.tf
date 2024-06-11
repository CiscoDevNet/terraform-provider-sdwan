resource "sdwan_dns_security_policy_definition" "example" {
  name                                      = "Example"
  description                               = "Example"
  domain_list_id                            = "84f10c9d-def7-45a3-8c64-6df26163c861"
  local_domain_bypass_enabled               = false
  match_all_vpn                             = true
  dnscrypt                                  = true
  umbrella_dns_default                      = true
  cisco_sig_credentials_feature_template_id = "885c3bce-eab6-4d45-8ff3-488f0df4072a"
}
