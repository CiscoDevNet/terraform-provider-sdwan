resource "sdwan_cisco_trustsec_feature_template" "example" {
  name                      = "Example"
  description               = "My Example"
  device_types              = ["vedge-C8000V"]
  device_sgt                = 100
  credentials_id            = "example"
  credentials_password      = "MyPassword"
  enable_enforcement        = true
  enable_sxp                = true
  sxp_source_ip             = "1.2.3.4"
  sxp_default_password      = "MyPassword"
  sxp_key_chain             = "example"
  sxp_log_binding_changes   = false
  sxp_reconciliation_period = 120
  sxp_retry_period          = 120
  speaker_hold_time         = 120
  listener_hold_time_min    = 90
  listener_hold_time_max    = 180
  sxp_node_id_type          = "interface-name"
  sxp_node_id               = "VirtualPortGroup"
  sxp_connection_list = [
    {
      connection_peer_ip       = "1.2.3.4"
      connection_source_ip     = "2.3.4.5"
      connection_preshared_key = "default"
      connection_mode          = "local"
      connection_mode_type     = "listener"
      connection_min_hold_time = 0
      connection_max_hold_time = 0
      connection_vpn_id        = 0
    }
  ]
}
