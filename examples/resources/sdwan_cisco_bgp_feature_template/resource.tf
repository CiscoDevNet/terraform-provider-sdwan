resource "sdwan_cisco_bgp_feature_template" "example" {
  name                = "Example"
  description         = "My Example"
  device_types        = ["vedge-C8000V"]
  as_number           = "65000"
  shutdown            = true
  router_id           = "1.2.3.4"
  propagate_aspath    = true
  propagate_community = true
  ipv4_route_targets = [
    {
      vpn_id = 1
      export = [
        {
          asn_ip = "10:100"
        }
      ]
      import = [
        {
          asn_ip = "10:100"
        }
      ]
    }
  ]
  ipv6_route_targets = [
    {
      vpn_id = 1
      export = [
        {
          asn_ip = "10:100"
        }
      ]
      import = [
        {
          asn_ip = "10:100"
        }
      ]
    }
  ]
  mpls_interfaces = [
    {
      interface_name = "GigabitEthernet0"
    }
  ]
  distance_external  = 30
  distance_internal  = 210
  distance_local     = 30
  keepalive          = 90
  holdtime           = 220
  always_compare_med = true
  deterministic_med  = true
  missing_med_worst  = true
  compare_router_id  = true
  multipath_relax    = true
  address_families = [
    {
      family_type = "ipv4-unicast"
      ipv4_aggregate_addresses = [
        {
          prefix       = "10.0.0.0/8"
          as_set_path  = true
          summary_only = true
        }
      ]
      ipv6_aggregate_addresses = [
        {
          prefix       = "10.0.0.0/8"
          as_set_path  = true
          summary_only = true
        }
      ]
      ipv4_networks = [
        {
          prefix = "10.2.2.0/24"
        }
      ]
      ipv6_networks = [
        {
          prefix = "10.2.2.0/24"
        }
      ]
      maximum_paths                 = 8
      default_information_originate = true
      table_map_policy              = "MAP1"
      table_map_filter              = true
      redistribute_routes = [
        {
          protocol     = "ospf"
          route_policy = "POLICY1"
        }
      ]
    }
  ]
  ipv4_neighbors = [
    {
      address             = "10.2.2.2"
      description         = "My neighbor"
      shutdown            = true
      remote_as           = "65001"
      keepalive           = 30
      holdtime            = 90
      source_interface    = "GigabitEthernet1"
      next_hop_self       = true
      send_community      = false
      send_ext_community  = false
      ebgp_multihop       = 10
      password            = "cisco123"
      send_label          = true
      send_label_explicit = true
      as_override         = true
      allow_as_in         = 5
      address_families = [
        {
          family_type                   = "ipv4-unicast"
          maximum_prefixes              = 10000
          maximum_prefixes_threshold    = 80
          maximum_prefixes_restart      = 180
          maximum_prefixes_warning_only = true
          route_policies = [
            {
              direction   = "in"
              policy_name = "POLICY1"
            }
          ]
        }
      ]
    }
  ]
  ipv6_neighbors = [
    {
      address             = "2001:1::1"
      description         = "My neighbor"
      shutdown            = true
      remote_as           = "65001"
      keepalive           = 30
      holdtime            = 90
      source_interface    = "GigabitEthernet1"
      next_hop_self       = true
      send_community      = false
      send_ext_community  = false
      ebgp_multihop       = 10
      password            = "cisco123"
      send_label          = true
      send_label_explicit = true
      as_override         = true
      allow_as_in         = 5
      address_families = [
        {
          family_type                   = "ipv6-unicast"
          maximum_prefixes              = 10000
          maximum_prefixes_threshold    = 80
          maximum_prefixes_restart      = 180
          maximum_prefixes_warning_only = true
          route_policies = [
            {
              direction   = "in"
              policy_name = "POLICY1"
            }
          ]
        }
      ]
    }
  ]
}
