resource "sdwan_localized_policy" "example" {
  name                          = "Example"
  description                   = "My description"
  flow_visibility_ipv4          = true
  flow_visibility_ipv6          = true
  application_visibility_ipv4   = true
  application_visibility_ipv6   = true
  cloud_qos                     = true
  cloud_qos_service_side        = true
  implicit_acl_logging          = true
  log_frequency                 = 1000
  ipv4_visibility_cache_entries = 1000
  ipv6_visibility_cache_entries = 1000
  definitions = [
    {
      id   = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
      type = "acl"
    }
  ]
}
