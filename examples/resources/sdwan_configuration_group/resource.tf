resource "sdwan_configuration_group" "example" {
  name        = "CG_1"
  description = "My config group 1"
  solution    = "sdwan"
  feature_profiles = [
    {
      id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
    }
  ]
}
