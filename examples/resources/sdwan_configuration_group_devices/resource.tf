resource "sdwan_configuration_group_devices" "example" {
  configuration_group_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  devices = [
    {
      id = "C8K-15411CCC-D476-0B3B-21F2-5D6AC387EE7B"
    }
  ]
}
