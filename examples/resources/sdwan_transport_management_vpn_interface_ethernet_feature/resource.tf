resource "sdwan_transport_management_vpn_interface_ethernet_feature" "example" {
  name                                = "Example"
  description                         = "My Example"
  feature_profile_id                  = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  transport_management_vpn_feature_id = "140331f6-5418-4755-a059-13c77eb96037"
  shutdown                            = true
  interface_name                      = "GigabitEthernet1"
  interface_description               = "Transport Management VPN Interface Ethernet"
  ipv4_address                        = "1.2.3.4"
  ipv4_subnet_mask                    = "0.0.0.0"
  ipv4_secondary_addresses = [
    {
      address     = "1.2.3.4"
      subnet_mask = "0.0.0.0"
    }
  ]
  ipv4_dhcp_helper           = ["1.2.3.4"]
  ipv4_iperf_server          = "example"
  ipv4_auto_detect_bandwidth = false
  ipv6_address               = "2001:0:0:1::/64"
  arp_entries = [
    {
      ip_address  = "1.2.3.4"
      mac_address = "00-B0-D0-63-C2-26"
    }
  ]
  duplex                = "full"
  mac_address           = "00-B0-D0-63-C2-26"
  ip_mtu                = 1500
  interface_mtu         = 1500
  tcp_mss               = 505
  speed                 = "2500"
  arp_timeout           = 1200
  autonegotiate         = false
  media_type            = "rj45"
  load_interval         = 30
  icmp_redirect_disable = true
  ip_directed_broadcast = false
}
