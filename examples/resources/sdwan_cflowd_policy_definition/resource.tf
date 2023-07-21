resource "sdwan_cflowd_policy_definition" "example" {
  name                  = "Example"
  description           = "My description"
  active_flow_timeout   = 100
  inactive_flow_timeout = 10
  sampling_interval     = 10
  flow_refresh          = 120
  protocol              = "ipv4"
  tos                   = true
  remarked_dscp         = true
  collectors = [
    {
      vpn_id           = 1
      ip_address       = "10.0.0.1"
      port             = 12345
      transport        = "transport_tcp"
      source_interface = "Ethernet1"
      export_spreading = "enable"
    }
  ]
}
