resource "sdwan_configuration_group" "example" {
  name        = "CG_1"
  description = "My config group 1"
  solution    = "sdwan"
  feature_profiles = [
    {
      id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
    }
  ]
  devices = [
    {
      id = "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"
      variables = [
        {
          name  = "host_name"
          value = "edge1"
        }
      ]
    }
  ]
}
