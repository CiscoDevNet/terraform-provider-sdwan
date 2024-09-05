resource "sdwan_transport_routing_bgp_feature" "example" {
  name                     = "Example"
  description              = "My Example"
  feature_profile_id       = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  as_number                = 429
  router_id                = "1.2.3.4"
  propagate_as_path        = false
  propagate_community      = false
  external_routes_distance = 20
  internal_routes_distance = 200
  local_routes_distance    = 20
  keepalive_time           = 60
  hold_time                = 180
  always_compare_med       = false
  deterministic_med        = false
  missing_med_as_worst     = false
  compare_router_id        = false
  multipath_relax          = false
  ipv4_neighbors = [
    {
      address                 = "1.2.3.4"
      description             = "neighbor1"
      shutdown                = false
      remote_as               = 200
      local_as                = 200
      keepalive_time          = 40
      hold_time               = 200
      update_source_interface = "GigabitEthernet0"
      next_hop_self           = false
      send_community          = true
      send_extended_community = true
      ebgp_multihop           = 1
      password                = "myPassword"
      send_label              = true
      explicit_null           = false
      as_override             = false
      allowas_in_number       = 1
      address_families = [
        {
          family_type            = "ipv4-unicast"
          max_number_of_prefixes = 2000
          threshold              = 75
          policy_type            = "restart"
          restart_interval       = 30
        }
      ]
    }
  ]
  ipv6_neighbors = [
    {
      address                 = "2001::1"
      description             = "neighbor2"
      shutdown                = false
      remote_as               = 200
      local_as                = 200
      keepalive_time          = 180
      hold_time               = 60
      update_source_interface = "Loopback1"
      next_hop_self           = true
      send_community          = true
      send_extended_community = true
      ebgp_multihop           = 3
      password                = "myPassword"
      as_override             = true
      allowas_in_number       = 3
      address_families = [
        {
          family_type            = "ipv6-unicast"
          max_number_of_prefixes = 2000
          threshold              = 75
          policy_type            = "restart"
          restart_interval       = 30
        }
      ]
    }
  ]
  ipv4_aggregate_addresses = [
    {
      network_address = "10.10.0.0"
      subnet_mask     = "255.255.0.0"
      as_set_path     = false
      summary_only    = false
    }
  ]
  ipv4_networks = [
    {
      network_address = "10.10.0.0"
      subnet_mask     = "255.255.0.0"
    }
  ]
  ipv4_eibgp_maximum_paths = 1
  ipv4_originate           = false
  ipv4_table_map_filter    = false
  ipv6_aggregate_addresses = [
    {
      aggregate_prefix = "3001::1/128"
      as_set_path      = false
      summary_only     = false
    }
  ]
  ipv6_networks = [
    {
      network_prefix = "2001:0DB8:0000:000b::/64"
    }
  ]
  ipv6_eibgp_maximum_paths = 2
  ipv6_originate           = true
  ipv6_table_map_filter    = false
  mpls_interfaces = [
    {
      interface_name = "GigabitEthernet1"
    }
  ]
}
