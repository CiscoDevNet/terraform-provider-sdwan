resource "sdwan_transport_cellular_controller_feature" "example" {
  name                 = "Example"
  description          = "My Example"
  feature_profile_id   = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  cellular_id          = "1"
  primary_sim_slot     = 0
  sim_failover_retries = 5
  sim_failover_timeout = 3
  firmware_auto_sim    = true
}
