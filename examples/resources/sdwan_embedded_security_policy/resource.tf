resource "sdwan_embedded_security_policy" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  assembly = [
    {
      entries = [
        {
        }
      ]
    }
  ]
  tcp_syn_flood_limit             = "432"
  max_incomplete_tcp_limit        = "12345"
  max_incomplete_udp_limit        = "12345"
  max_incomplete_icmp_limit       = "12345"
  audit_trail                     = "on"
  unified_logging                 = "off"
  session_reclassify_allow        = "off"
  imcp_unreachable_allow          = "off"
  failure_mode                    = "close"
  nat                             = true
  download_url_database_on_device = false
  resource_profile                = "low"
}
