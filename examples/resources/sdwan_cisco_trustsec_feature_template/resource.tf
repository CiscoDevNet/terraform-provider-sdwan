resource "sdwan_cisco_trustsec_feature_template" "example" {
  name                       = "Example"
  description                = "My Example"
  device_types               = ["vedge-C8000V"]
  device_sgt                 = 100
  credentials_id             = "user1"
  credentials_password       = "MyPassword"
  enable_enforcement         = true
  enable_sxp                 = true
  sxp_source_ip              = "1.2.3.4"
  sxp_default_password       = "MyPassword"
  sxp_key_chain              = "keychain1"
  sxp_log_binding_changes    = false
  sxp_reconciliation_period  = 120
  sxp_retry_period           = 120
  speaker_hold_time          = 120
  minimum_listener_hold_time = 90
  maximum_listener_hold_time = 180
  sxp_node_id_type           = "interface-name"
  sxp_node_id                = "VirtualPortGroup"
  sxp_connections = [
    {
      peer_ip           = "1.2.3.4"
      source_ip         = "2.3.4.5"
      preshared_key     = "default"
      mode              = "local"
      mode_type         = "listener"
      minimum_hold_time = 100
      maximum_hold_time = 200
      vpn_id            = 0
    }
  ]
}
