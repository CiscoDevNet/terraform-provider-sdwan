resource "sdwan_transport_management_vpn_profile_parcel" "example" {
  name                            = "Example"
  description                     = "My Example"
  feature_profile_id              = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  basic_configuration_description = "example"
  primary_dns_address_ipv4        = "1.2.3.4"
  secondary_dns_address_ipv4      = "2.3.4.5"
  primary_dns_address_ipv6        = "2001:0:0:1::0"
  secondary_dns_address_ipv6      = "2001:0:0:2::0"
  host_mappings = [
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
      ipv4_route_gateway_next_ho = [
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
          address                 = "2001:0:0:1::/64"
          administrative_distance = 1
        }
      ]
    }
  ]
}
