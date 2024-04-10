resource "sdwan_system_basic_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  timezone           = "UTC"
  config_description = "example"
  location           = "example"
  longitude          = -77
  latitude           = 38
  enable             = true
  range              = 100
  sms_enable         = true
  mobile_numbers = [
    {
      number = "+11111233"
    }
  ]
  device_groups             = ["example"]
  controller_group_list     = [1]
  overlay_id                = 1
  port_offset               = 19
  port_hop                  = true
  control_session_pps       = 300
  track_transport           = true
  track_interface_tag       = 2
  console_baud_rate         = "4800"
  max_omp_sessions          = 24
  multi_tenant              = false
  track_default_gateway     = true
  admin_tech_on_failure     = true
  idle_timeout              = 10
  on_demand_enable          = true
  on_demand_idle_timeout    = 10
  transport_gateway         = false
  epfr                      = "disabled"
  site_type                 = ["type-1"]
  affinity_group_number     = 1
  affinity_group_preference = [1]
  affinity_preference_auto  = false
  affinity_per_vrf = [
    {
      affinity_group_number = 1
      vrf_range             = "123-456"
    }
  ]
}
