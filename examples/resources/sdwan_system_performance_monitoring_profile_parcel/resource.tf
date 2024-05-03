resource "sdwan_system_performance_monitoring_profile_parcel" "example" {
  name                        = "Example"
  description                 = "My Example"
  feature_profile_id          = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  app_perf_monitor_enabled    = true
  app_perf_monitor_app_group  = ["amazon-group"]
  monitoring_config_enabled   = true
  monitoring_config_interval  = "30"
  event_driven_config_enabled = true
  event_driven_events         = ["SLA_CHANGE"]
}
