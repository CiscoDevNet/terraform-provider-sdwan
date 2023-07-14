resource "sdwan_cisco_ospf_feature_template" "example" {
  name                                      = "Example"
  description                               = "My Example"
  device_types                              = ["vedge-C8000V"]
  router_id                                 = "1.2.3.4"
  auto_cost_reference_bandwidth             = 100000
  compatible_rfc1583                        = true
  default_information_originate             = true
  default_information_originate_always      = true
  default_information_originate_metric      = 100
  default_information_originate_metric_type = "type1"
  distance_external                         = 111
  distance_inter_area                       = 111
  distance_intra_area                       = 112
  timers_spf_delay                          = 300
  timers_spf_initial_holf                   = 2000
  timers_spf_max_hold                       = 20000
  redistribute = [
    {
      protocol     = "static"
      route_policy = "RP1"
      nat_dia      = true
    }
  ]
  max_metric_router_lsa = [
    {
      ad_type = "on-startup"
      time    = 100
    }
  ]
  route_policies = [
    {
      direction   = "in"
      policy_name = "POLICY1"
    }
  ]
  areas = [
    {
      area_number     = 1
      stub_no_summary = false
      nssa_no_summary = true
      interfaces = [
        {
          name                                 = "e1"
          hello_interval                       = 20
          dead_interval                        = 60
          retransmit_interval                  = 10
          cost                                 = 100
          priority                             = 10
          network                              = "point-to-point"
          passive_interface                    = true
          authentication_type                  = "message-digest"
          authentication_message_digest_key_id = 1
          authentication_message_digest_key    = "cisco123"
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
