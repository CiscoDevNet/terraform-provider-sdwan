resource "sdwan_cisco_ospfv3_feature_template" "example" {
  name                                           = "Example"
  description                                    = "My Example"
  device_types                                   = ["vedge-C8000V"]
  router_id_ipv4                                 = "1.2.3.4"
  auto_cost_reference_bandwidth_ipv4             = 100000
  compatible_rfc1583_ipv4                        = true
  default_information_originate_ipv4             = true
  default_information_originate_always_ipv4      = true
  default_information_originate_metric_ipv4      = 100
  default_information_originate_metric_type_ipv4 = "type1"
  distance_external_ipv4                         = 111
  distance_inter_area_ipv4                       = 111
  distance_intra_area_ipv4                       = 112
  timers_spf_delay_ipv4                          = 300
  timers_spf_initial_hold_ipv4                   = 2000
  timers_spf_max_hold_ipv4                       = 20000
  distance_ipv4                                  = 110
  policy_name_ipv4                               = "example"
  filter_ipv4                                    = false
  redistribute_ipv4 = [
    {
      protocol     = "static"
      route_policy = "RP1"
      nat_dia      = true
    }
  ]
  max_metric_router_lsa_ipv4 = [
    {
      ad_type = "on-startup"
      time    = 100
    }
  ]
  areas_ipv4 = [
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
          authentication_type = "message-digest"
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
  router_id_ipv6                                 = "1.2.3.4"
  auto_cost_reference_bandwidth_ipv6             = 100000
  compatible_rfc1583_ipv6                        = true
  default_information_originate_ipv6             = true
  default_information_originate_always_ipv6      = true
  default_information_originate_metric_ipv6      = 100
  default_information_originate_metric_type_ipv6 = "type1"
  distance_external_ipv6                         = 111
  distance_inter_area_ipv6                       = 111
  distance_intra_area_ipv6                       = 112
  timers_spf_delay_ipv6                          = 300
  timers_spf_initial_hold_ipv6                   = 2000
  timers_spf_max_hold_ipv6                       = 20000
  distance_ipv6                                  = 110
  policy_name_ipv6                               = "example"
  filter_ipv6                                    = false
  redistribute_ipv6 = [
    {
      protocol     = "static"
      route_policy = "RP1"
    }
  ]
  max_metric_router_lsa_ipv6 = [
    {
      ad_type = "on-startup"
      time    = 100
    }
  ]
  areas_ipv6 = [
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
          authentication_type = "message-digest"
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
}
