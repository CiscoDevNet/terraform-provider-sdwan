resource "sdwan_cellular_controller_feature_template" "example" {
  name                  = "Example"
  description           = "My Example"
  device_types          = ["vedge-C8000V"]
  cellular_interface_id = "1"
  data_profiles = [
    {
      slot_number    = 1
      data_profile   = 8
      attach_profile = 8
    }
  ]
  primary_sim_slot     = 100
  sim_failover_retries = 160
  sim_failover_timeout = 3
  firmware_auto_sim    = false
}
