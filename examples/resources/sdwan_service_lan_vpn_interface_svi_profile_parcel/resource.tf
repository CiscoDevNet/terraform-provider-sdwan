resource "sdwan_service_lan_vpn_interface_svi_feature" "example" {
  name                       = "Example"
  description                = "My Example"
  feature_profile_id         = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  service_lan_vpn_feature_id = "140331f6-5418-4755-a059-13c77eb96037"
  shutdown                   = false
  interface_name             = "Vlan1"
  interface_description      = "SVI"
  interface_mtu              = 1500
  ip_mtu                     = 1500
  ipv4_address               = "1.2.3.4"
  ipv4_subnet_mask           = "0.0.0.0"
  ipv4_secondary_addresses = [
    {
      address          = "2.3.4.5"
      ipv4_subnet_mask = "0.0.0.0"
    }
  ]
  ipv4_dhcp_helpers = ["4.5.6.7"]
  ipv6_address      = "2001:0:0:1::0/32"
  ipv6_secondary_addresses = [
    {
      address = "::2/32"
    }
  ]
  ipv6_dhcp_helpers = [
    {
      address = "2001:0:0:1::0"
      vpn     = 1
    }
  ]
  arps = [
    {
      ip_address  = "1.2.3.4"
      mac_address = "00-B0-D0-63-C2-26"
    }
  ]
  ipv4_vrrps = [
    {
      group_id    = 1
      priority    = 100
      timer       = 1000
      track_omp   = false
      prefix_list = "prefix"
      address     = "1.2.3.4"
      secondary_addresses = [
        {
          address = "2.3.4.5"
        }
      ]
      tloc_prefix_change       = true
      tloc_prefix_change_value = 100
    }
  ]
  ipv6_vrrps = [
    {
      group_id          = 1
      priority          = 100
      timer             = 1000
      track_omp         = false
      track_prefix_list = "1"
      addresses = [
        {
          link_local_address = "1::1"
          global_address     = "1::1/24"
        }
      ]
      secondary_addresses = [
        {
          prefix = "::20/32"
        }
      ]
    }
  ]
  enable_dhcpv6         = false
  tcp_mss               = 1024
  arp_timeout           = 1200
  ip_directed_broadcast = false
  icmp_redirect_disable = true
}
