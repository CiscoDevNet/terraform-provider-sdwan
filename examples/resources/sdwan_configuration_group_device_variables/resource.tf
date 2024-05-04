resource "sdwan_configuration_group_device_variables" "example" {
  configuration_group_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  solution               = "sdwan"
  devices = [
    {
      device_id = "C8K-15411CCC-D476-0B3B-21F2-5D6AC387EE7B"
      variables = [
        {
          name  = "host_name"
          value = "edge1"
        }
      ]
    }
  ]
}
