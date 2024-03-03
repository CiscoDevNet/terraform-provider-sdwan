resource "sdwan_cisco_system_feature_template" "example" {
  name               = "Example"
  description        = "My Example"
  device_types       = ["vedge-C8000V"]
  timezone           = "UTC"
  hostname           = "Router1"
  system_description = "My Description"
  location           = "Building 1"
  latitude           = 40
  longitude          = 50
  geo_fencing        = true
  geo_fencing_range  = 1000
  geo_fencing_sms    = true
  geo_fencing_sms_phone_numbers = [
    {
      number = "+1234567"
    }
  ]
  device_groups         = ["group1"]
  controller_group_list = [1]
  system_ip             = "5.5.5.5"
  overlay_id            = 1
  site_id               = 1
  port_offset           = 1
  port_hopping          = true
  control_session_pps   = 300
  track_transport       = true
  track_interface_tag   = 1
  console_baud_rate     = "115200"
  max_omp_sessions      = 5
  multi_tenant          = true
  track_default_gateway = true
  admin_tech_on_failure = true
  idle_timeout          = 100
  trackers = [
    {
      name        = "tracker1"
      endpoint_ip = "5.6.7.8"
      threshold   = 300
      interval    = 60
      multiplier  = 3
      type        = "interface"
    }
  ]
  object_trackers = [
    {
      object_number = 1
      interface     = "e1"
    }
  ]
  on_demand_tunnel              = true
  on_demand_tunnel_idle_timeout = 10
  affinity_group_number         = 5
  affinity_group_preference     = [1]
  transport_gateway             = true
  enable_mrf_migration          = "enabled"
  migration_bgp_community       = 100
}
