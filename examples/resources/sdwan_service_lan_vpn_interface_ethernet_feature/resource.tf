resource "sdwan_service_lan_vpn_interface_ethernet_feature" "example" {
  name                       = "Example"
  description                = "My Example"
  feature_profile_id         = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  service_lan_vpn_feature_id = "140331f6-5418-4755-a059-13c77eb96037"
  shutdown                   = false
  interface_name             = "GigabitEthernet3"
  interface_description      = "LAN"
  ipv4_address               = "1.2.3.4"
  ipv4_subnet_mask           = "0.0.0.0"
  ipv4_secondary_addresses = [
    {
      address     = "1.2.3.5"
      subnet_mask = "0.0.0.0"
    }
  ]
  ipv4_dhcp_helper = ["1.2.3.4"]
  ipv6_dhcp_helpers = [
    {
      address           = "2001:0:0:1::0"
      dhcpv6_helper_vpn = 1
    }
  ]
  ipv4_nat               = false
  ipv4_nat_type          = "pool"
  ipv4_nat_range_start   = "1.2.3.4"
  ipv4_nat_range_end     = "4.5.6.7"
  ipv4_nat_prefix_length = 1
  ipv4_nat_overload      = true
  ipv4_nat_loopback      = "123"
  ipv4_nat_udp_timeout   = 123
  ipv4_nat_tcp_timeout   = 123
  static_nats = [
    {
      source_ip    = "1.2.3.4"
      translate_ip = "2.3.4.5"
      direction    = "inside"
      source_vpn   = 0
    }
  ]
  ipv6_nat         = true
  nat64            = false
  acl_shaping_rate = 12
  ipv6_vrrps = [
    {
      group_id  = 1
      priority  = 100
      timer     = 1000
      track_omp = false
      ipv6_addresses = [
        {
          link_local_address = "1::1"
          global_address     = "1::1/24"
        }
      ]
    }
  ]
  ipv4_vrrps = [
    {
      group_id  = 1
      priority  = 100
      timer     = 1000
      track_omp = false
      address   = "1.2.3.4"
      secondary_addresses = [
        {
          address     = "2.3.4.5"
          subnet_mask = "0.0.0.0"
        }
      ]
      tloc_prefix_change     = true
      tloc_pref_change_value = 100
      tracking_objects = [
        {
          tracker_id      = "1b270f6d-479b-47e3-ab0b-51bc6811a303"
          tracker_action  = "Decrement"
          decrement_value = 100
        }
      ]
    }
  ]
  arps = [
    {
      ip_address  = "1.2.3.4"
      mac_address = "00-B0-D0-63-C2-26"
    }
  ]
  trustsec_enable_sgt_propogation      = false
  trustsec_propogate                   = true
  trustsec_security_group_tag          = 123
  trustsec_enable_enforced_propogation = false
  trustsec_enforced_security_group_tag = 1234
  duplex                               = "full"
  mac_address                          = "00-B0-D0-63-C2-26"
  ip_mtu                               = 1500
  interface_mtu                        = 1500
  tcp_mss                              = 500
  speed                                = "1000"
  arp_timeout                          = 1200
  autonegotiate                        = false
  media_type                           = "auto-select"
  load_interval                        = 30
  tracker                              = "TRACKER1"
  icmp_redirect_disable                = true
  xconnect                             = "1"
  ip_directed_broadcast                = false
}
