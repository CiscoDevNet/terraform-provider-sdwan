resource "sdwan_service_routing_ospf_profile_parcel" "example" {
  name                           = "Example"
  description                    = "My Example"
  feature_profile_id             = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  router_id                      = "1.2.3.4"
  reference_bandwidth            = 101
  rfc_1583_compatible            = true
  originate                      = false
  always                         = false
  metric                         = 1
  metric_type                    = "type1"
  distance_for_external_routes   = 110
  distance_for_inter_area_routes = 110
  distance_for_intra_area_routes = 110
  spf_calculation_deplay         = 200
  initial_hold_time              = 1000
  maximum_hold_time              = 10000
  redistributes = [
    {
      protocol = "static"
      dia      = true
    }
  ]
  router_lsas = [
    {
      ad_type = "on-startup"
      time    = 5
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
          interface_cost             = 10
          designated_router_priority = 1
          ospf_network_type          = "broadcast"
          passive_interface          = false
          type                       = "message-digest"
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
