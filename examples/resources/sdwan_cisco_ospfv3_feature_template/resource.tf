resource "sdwan_cisco_ospfv3_feature_template" "example" {
  name                                           = "Example"
  description                                    = "My Example"
  device_types                                   = ["vedge-C8000V"]
  ipv4_router_id                                 = "1.2.3.4"
  ipv4_auto_cost_reference_bandwidth             = 100000
  ipv4_compatible_rfc1583                        = true
  ipv4_default_information_originate             = true
  ipv4_default_information_originate_always      = true
  ipv4_default_information_originate_metric      = 100
  ipv4_default_information_originate_metric_type = "type1"
  ipv4_distance_external                         = 111
  ipv4_distance_inter_area                       = 111
  ipv4_distance_intra_area                       = 112
  ipv4_timers_spf_delay                          = 300
  ipv4_timers_spf_initial_hold                   = 2000
  ipv4_timers_spf_max_hold                       = 20000
  ipv4_distance                                  = 110
  ipv4_policy_name                               = "POLICY1"
  ipv4_filter                                    = false
  ipv4_redistributes = [
    {
      protocol     = "static"
      route_policy = "RP1"
      nat_dia      = true
    }
  ]
  ipv4_max_metric_router_lsas = [
    {
      ad_type = "on-startup"
      time    = 100
    }
  ]
  ipv4_areas = [
    {
      area_number     = 1
      stub            = false
      stub_no_summary = false
      nssa            = false
      nssa_no_summary = true
      translate       = "always"
      normal          = false
      interfaces = [
        {
          name                = "e1"
          hello_interval      = 20
          dead_interval       = 60
          retransmit_interval = 10
          cost                = 100
          network             = "point-to-point"
          passive_interface   = true
          authentication_type = "md5"
          authentication_key  = "authenticationKey"
          ipsec_spi           = 256
        }
      ]
      ranges = [
        {
          address      = "1.1.1.0/24"
          cost         = 100
          no_advertise = true
        }
      ]
    }
  ]
  ipv6_router_id                                 = "1.2.3.4"
  ipv6_auto_cost_reference_bandwidth             = 100000
  ipv6_compatible_rfc1583                        = true
  ipv6_default_information_originate             = true
  ipv6_default_information_originate_always      = true
  ipv6_default_information_originate_metric      = 100
  ipv6_default_information_originate_metric_type = "type1"
  ipv6_distance_external                         = 111
  ipv6_distance_inter_area                       = 111
  ipv6_distance_intra_area                       = 112
  ipv6_timers_spf_delay                          = 300
  ipv6_timers_spf_initial_hold                   = 2000
  ipv6_timers_spf_max_hold                       = 20000
  ipv6_distance                                  = 110
  ipv6_policy_name                               = "POLICY2"
  ipv6_filter                                    = false
  ipv6_redistributes = [
    {
      protocol     = "static"
      route_policy = "RP1"
    }
  ]
  ipv6_max_metric_router_lsas = [
    {
      ad_type = "on-startup"
      time    = 100
    }
  ]
  ipv6_areas = [
    {
      area_number     = 1
      stub            = false
      stub_no_summary = false
      nssa            = false
      nssa_no_summary = true
      translate       = "always"
      normal          = false
      interfaces = [
        {
          name                = "e1"
          hello_interval      = 20
          dead_interval       = 60
          retransmit_interval = 10
          cost                = 100
          network             = "point-to-point"
          passive_interface   = true
          authentication_type = "md5"
          authentication_key  = "authenticationKey"
          ipsec_spi           = 256
        }
      ]
      ranges = [
        {
          address      = "2001::/48"
          cost         = 100
          no_advertise = true
        }
      ]
    }
  ]
}
