resource "sdwan_service_routing_ospf_profile_parcel" "example" {
  name                                      = "Example"
  description                               = "My Example"
  feature_profile_id                        = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  router_id                                 = "1.2.3.4"
  reference_bandwidth                       = 101
  rfc_1583_compatible                       = true
  default_information_originate             = false
  default_information_originate_always      = false
  default_information_originate_metric      = 1
  default_information_originate_metric_type = "type1"
  distance_external                         = 110
  distance_inter_area                       = 110
  distance_intra_area                       = 110
  spf_calculation_delay                     = 200
  spf_initial_hold_time                     = 1000
  spf_maximum_hold_time                     = 10000
  redistributes = [
    {
      protocol = "static"
      nat_dia  = true
    }
  ]
  router_lsas = [
    {
      type = "on-startup"
      time = 5
    }
  ]
  areas = [
    {
      area_number = 1
      area_type   = "stub"
      no_summary  = false
      interfaces = [
        {
          name                       = "GigabitEthernet2"
          hello_interval             = 10
          dead_interval              = 40
          lsa_retransmit_interval    = 5
          cost                       = 10
          designated_router_priority = 1
          network_type               = "broadcast"
          passive_interface          = false
          authentication_type        = "message-digest"
          message_digest_key_id      = 7
          message_digest_key         = "sdjfhsghbjdjr"
        }
      ]
      ranges = [
        {
          ip_address   = "10.1.1.0"
          subnet_mask  = "255.255.255.0"
          cost         = 1
          no_advertise = false
        }
      ]
    }
  ]
}
