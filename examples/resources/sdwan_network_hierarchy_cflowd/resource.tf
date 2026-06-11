resource "sdwan_network_hierarchy_cflowd" "example" {
  node_id                = "5333c142-394b-47a7-afa6-760c44ca3cb5"
  flow_active_timeout    = 600
  flow_inactive_timeout  = 60
  flow_refresh_time      = 600
  flow_sampling_interval = 1
  collect_tloc_loopback  = true
  protocol               = "ipv4"
  collect_tos            = true
  collect_dscp_output    = true
  collectors = [
    {
      vpn_id             = 1
      address            = "10.0.0.1"
      udp_port           = 4739
      export_spread      = true
      bfd_metrics_export = true
      export_interval    = 600
    }
  ]
}
