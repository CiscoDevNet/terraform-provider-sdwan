resource "sdwan_vpn_interface_svi_feature_template" "example" {
  name                  = "Example"
  description           = "My Example"
  device_types          = ["vedge-C8000V"]
  if_name               = "Vlan100"
  interface_description = "VPN Interface SVI"
  ipv4_address          = "2.3.4.5"
  ipv4_secondary_addresses = [
    {
      ipv4_address = "4.5.6.7"
    }
  ]
  ipv6_address           = "2001:db8:85a3::8a2e:370:7334"
  ipv6_dhcp_client       = false
  ipv6_dhcp_distance     = 101
  ipv6_dhcp_rapid_commit = false
  ipv6_secondary_addresses = [
    {
      ipv6_address = "2001:db8:85a3::8a2e:370:7334"
    }
  ]
  ipv4_dhcp_helper = ["7.7.7.7"]
  ipv6_dhcp_helpers = [
    {
      address = "2001:db8:85a3::8a2e:370:7334"
      vpn_id  = 100
    }
  ]
  ip_directed_broadcast = true
  mtu                   = 1500
  ip_mtu                = 1500
  tcp_mss_adjust        = 1400
  shutdown              = false
  arp_timeout           = 100
  ipv4_access_lists = [
    {
      direction = "in"
      acl_name  = "ACL1"
    }
  ]
  ipv6_access_lists = [
    {
      direction = "in"
      acl_name  = "ACL2"
    }
  ]
  policers = [
    {
      direction    = "in"
      policer_name = "POLICER1"
    }
  ]
  static_arp_entries = [
    {
      ipv4_address = "3.4.4.5"
      mac_address  = "00:00:00:00:00:00"
    }
  ]
  ipv4_vrrps = [
    {
      group_id          = 1
      priority          = 100
      timer             = 1000
      track_omp         = true
      track_prefix_list = "TRACK1"
      ipv4_address      = "5.6.7.8"
      ipv4_secondary_addresses = [
        {
          ipv4_address = "8.8.8.8"
        }
      ]
      tloc_preference_change       = true
      tloc_preference_change_value = 100
      tracking_objects = [
        {
          name            = 100
          track_action    = "decrement"
          decrement_value = 10
        }
      ]
    }
  ]
  ipv6_vrrps = [
    {
      group_id          = 1
      priority          = 100
      timer             = 1000
      track_omp         = true
      track_prefix_list = "TRACK1"
      ipv6_addresses = [
        {
          link_local_address = "FE80::1/64"
          prefix             = "2001:db8:85a3::8a2e:370:7335"
        }
      ]
      ipv6_secondary_addresses = [
        {
          prefix = "2001:db8:85a3::8a2e:370:7336"
        }
      ]
    }
  ]
}
