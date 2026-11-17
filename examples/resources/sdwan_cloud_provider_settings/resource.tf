resource "sdwan_cloud_provider_settings" "example" {
  umbrella_org_id           = "123456"
  umbrella_auth_key_v2      = "aaaaaa"
  umbrella_auth_secret_v2   = "$CRYPT_CLUSTER$a8nXr4uYaCN66qTQ1737+A==$gjmYVxnKMqvUEjo3BfIGbg=="
  umbrella_sig_auth_key     = "bbbbbb"
  umbrella_sig_auth_secret  = "secret456"
  umbrella_dns_auth_key     = "cccccc"
  umbrella_dns_auth_secret  = "secret789"
  zscaler_organization      = "z223456"
  zscaler_partner_base_uri  = "zscaler.net"
  zscaler_partner_key       = "secret012"
  zscaler_username          = "user1"
  zscaler_password          = "pass123"
  cisco_sse_org_id          = "323456"
  cisco_sse_auth_key        = "33333444"
  cisco_sse_auth_secret     = "$CRYPT_CLUSTER$Gg4nVpFdldXga1hLKhdJrA==$hiFPirWJnqNxMq3l/m1ekw=="
  cisco_sse_context_sharing = true
}
