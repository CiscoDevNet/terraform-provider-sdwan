resource "sdwan_cloud_provider_settings" "example" {
  umbrella_org_id           = "123456"
  zscaler_organization      = "z223456"
  zscaler_partner_base_uri  = "zscaler.net"
  zscaler_username          = "user1"
  cisco_sse_org_id          = "323456"
  cisco_sse_context_sharing = true
}
