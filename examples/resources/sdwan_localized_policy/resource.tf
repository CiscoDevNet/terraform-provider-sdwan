resource "sdwan_localized_policy" "test" {
  name                          = "POLICY_1"
  description                   = "My localized policy"
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
  definitions = [{
    id   = sdwan_qos_map_policy_definition.qos_map.id
    type = sdwan_qos_map_policy_definition.qos_map.type
  }]
}
