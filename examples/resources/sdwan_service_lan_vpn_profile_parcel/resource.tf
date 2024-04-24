resource "sdwan_service_lan_vpn_profile_parcel" "example" {
  name                       = "Example"
  description                = "My Example"
  feature_profile_id         = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  vpn                        = 1
  config_description         = "VPN1"
  omp_admin_distance_ipv4    = 1
  omp_admin_distance_ipv6    = 1
  enable_sdwan_remote_access = false
  dns_primary_ipv4_address   = "1.2.3.4"
  dns_secondary_ipv4_address = "2.3.4.5"
  dns_primary_ipv6_address   = "2001:0:0:1::0"
  dns_secondary_ipv6_address = "2001:0:0:2::0"
  host_mappings = [
    {
      host_name  = "HOSTMAPPING1"
      list_of_ip = ["1.2.3.4"]
    }
  ]
  ipv4_static_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      next_hops = [
        {
          address                 = "1.2.3.4"
          administrative_distance = 1
        }
      ]
    }
  ]
  ipv6_routes = [
    {
      prefix = "2001:0:0:1::0/12"
      next_hops = [
        {
          address                 = "2001:0:0:1::0"
          administrative_distance = 1
        }
      ]
    }
  ]
  services = [
    {
      service_type   = "FW"
      ipv4_addresses = ["1.2.3.4"]
      tracking       = true
    }
  ]
  service_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      service         = "SIG"
      vpn             = 0
    }
  ]
  gre_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      interface       = ["gre01"]
      vpn             = 0
    }
  ]
  ipsec_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      interface       = ["ipsec01"]
    }
  ]
  nat_pools = [
    {
      nat_pool_name = 1
      prefix_length = 3
      range_start   = "1.2.3.4"
      range_end     = "2.3.4.5"
      overload      = true
      direction     = "inside"
    }
  ]
  nat_port_forwards = [
    {
      nat_pool_name        = 2
      source_port          = 122
      translate_port       = 330
      source_ip            = "1.2.3.4"
      translated_source_ip = "2.3.4.5"
      protocol             = "TCP"
    }
  ]
  static_nats = [
    {
      nat_pool_name        = 3
      source_ip            = "1.2.3.4"
      translated_source_ip = "2.3.4.5"
      static_nat_direction = "inside"
    }
  ]
  nat_64_v4_pool = [
    {
      nat64_v4_pool_name        = "NATPOOL1"
      nat64_v4_pool_range_start = "1.2.3.4"
      nat64_v4_pool_range_end   = "2.3.4.5"
      overload                  = false
    }
  ]
  ipv4_import_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
  ipv4_export_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
  ipv6_import_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
  ipv6_export_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
}
