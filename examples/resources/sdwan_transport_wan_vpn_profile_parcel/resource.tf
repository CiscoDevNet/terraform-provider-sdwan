resource "sdwan_transport_wan_vpn_feature" "example" {
  name                       = "Example"
  description                = "My Example"
  feature_profile_id         = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  vpn                        = 0
  enhance_ecmp_keying        = true
  primary_dns_address_ipv4   = "1.2.3.4"
  secondary_dns_address_ipv4 = "2.3.4.5"
  primary_dns_address_ipv6   = "2001:0:0:1::0"
  secondary_dns_address_ipv6 = "2001:0:0:2::0"
  new_host_mappings = [
    {
      host_name            = "example"
      list_of_ip_addresses = ["1.2.3.4"]
    }
  ]
  ipv4_static_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      gateway         = "nextHop"
      next_hops = [
        {
          address                 = "1.2.3.4"
          administrative_distance = 1
        }
      ]
      administrative_distance = 1
    }
  ]
  ipv6_static_routes = [
    {
      prefix = "2002::/16"
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
      service_type = "TE"
    }
  ]
  nat_64_v4_pools = [
    {
      nat64_v4_pool_name        = "example"
      nat64_v4_pool_range_start = "203.0.113.50"
      nat64_v4_pool_range_end   = "203.0.113.100"
      nat64_v4_pool_overload    = false
    }
  ]
}
