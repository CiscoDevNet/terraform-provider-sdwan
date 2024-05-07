resource "sdwan_application_priority_qos_policy_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  target_interface   = ["{{interface_var_1}}"]
  qos_schedulers = [
    {
      drops           = "tail-drop"
      queue           = "0"
      bandwidth       = "10"
      scheduling_type = "llq"
    }
  ]
}
