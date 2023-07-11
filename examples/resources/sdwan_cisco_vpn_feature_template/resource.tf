resource "sdwan_cisco_vpn_feature_template" "example" {
  name                    = "Example"
  description             = "My Example"
  device_types            = ["vedge-C8000V"]
  vpn_id                  = 1
  vpn_name                = "VPN1"
  tenant_vpn_id           = 1
  organization_name       = "org1"
  omp_admin_distance_ipv4 = 10
  omp_admin_distance_ipv6 = 10
  enhance_ecmp_keying     = true
  dns_ipv4_servers = [
    {
      address = "9.9.9.9"
      role    = "primary"
    }
  ]
  dns_ipv6_servers = [
    {
      address = "2001::9"
      role    = "primary"
    }
  ]
  dns_hosts = [
    {
      hostname = "abc1"
      ip       = ["7.7.7.7"]
    }
  ]
  services = [
    {
      service_types = "FW"
      address       = ["8.8.8.8"]
      interface     = "e1"
      track_enable  = true
    }
  ]
  ipv4_static_service_routes = [
    {
      prefix  = "2.2.2.0/24"
      vpn_id  = 2
      service = "sig"
    }
  ]
  ipv4_static_routes = [
    {
      prefix   = "3.3.3.0/24"
      null0    = false
      distance = 10
      vpn_id   = 5
      dhcp     = false
      next_hops = [
        {
          address  = "11.1.1.1"
          distance = 20
        }
      ]
      track_next_hops = [
        {
          address  = "12.1.1.1"
          distance = 20
          tracker  = "tracker1"
        }
      ]
    }
  ]
  ipv6_static_routes = [
    {
      prefix = "2001::/48"
      null0  = false
      vpn_id = 5
      nat    = "NAT64"
      next_hops = [
        {
          address  = "2001::11"
          distance = 20
        }
      ]
    }
  ]
  ipv4_static_gre_routes = [
    {
      prefix    = "3.3.3.0/24"
      vpn_id    = 2
      interface = ["e1"]
    }
  ]
  ipv4_static_ipsec_routes = [
    {
      prefix    = "4.4.4.0/24"
      vpn_id    = 2
      interface = ["e1"]
    }
  ]
  omp_advertise_ipv4_routes = [
    {
      protocol          = "bgp"
      route_policy      = "rp1"
      protocol_sub_type = ["external"]
      prefixes = [
        {
          prefix_entry   = "1.1.1.0/24"
          aggregate_only = true
        }
      ]
    }
  ]
  omp_advertise_ipv6_routes = [
    {
      protocol          = "bgp"
      route_policy      = "rp1"
      protocol_sub_type = ["external"]
      prefixes = [
        {
          prefix_entry   = "2001:2::/48"
          aggregate_only = true
        }
      ]
    }
  ]
  nat64_pools = [
    {
      name                      = "POOL1"
      start_address             = "100.1.1.1"
      end_address               = "100.1.2.255"
      overload                  = true
      leak_from_global          = true
      leak_from_global_protocol = "rip"
      leak_to_global            = true
    }
  ]
  nat_pools = [
    {
      name          = 1
      prefix_length = 24
      range_start   = "101.1.1.1"
      range_end     = "101.1.2.255"
      overload      = true
      direction     = "inside"
      tracker_id    = 10
    }
  ]
  static_nat_rules = [
    {
      pool_name            = "POOL1"
      source_ip            = "10.1.1.1"
      translate_ip         = "105.1.1.1"
      static_nat_direction = "inside"
      tracker_id           = 10
    }
  ]
  static_nat_subnet_rules = [
    {
      source_ip_subnet     = "10.2.1.0"
      translate_ip_subnet  = "105.2.1.0"
      prefix_length        = 24
      static_nat_direction = "inside"
      tracker_id           = 10
    }
  ]
  port_forward_rules = [
    {
      pool_name      = "POOL2"
      source_port    = 5000
      translate_port = 6000
      source_ip      = "10.3.1.1"
      translate_ip   = "120.3.1.1"
      protocol       = "tcp"
    }
  ]
  route_global_imports = [
    {
      protocol          = "ospf"
      protocol_sub_type = ["external"]
      route_policy      = "policy1"
      redistributes = [
        {
          protocol     = "bgp"
          route_policy = "policy1"
        }
      ]
    }
  ]
  route_vpn_imports = [
    {
      source_vpn_id     = 5
      protocol          = "ospf"
      protocol_sub_type = ["external"]
      route_policy      = "policy1"
      redistributes = [
        {
          protocol     = "bgp"
          route_policy = "policy1"
        }
      ]
    }
  ]
  route_global_exports = [
    {
      protocol          = "ospf"
      protocol_sub_type = ["external"]
      route_policy      = "policy1"
      redistributes = [
        {
          protocol     = "bgp"
          route_policy = "policy1"
        }
      ]
    }
  ]
}
