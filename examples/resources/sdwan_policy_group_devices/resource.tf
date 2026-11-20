resource "sdwan_policy_group_devices" "example" {
  policy_group_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  solution        = "sdwan"
  devices = {
    "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B" = {
      deploy = true
      variables = {
        qos_interfaces = {
          list_value = ["GigabitEthernet1", "GigabitEthernet2"]
        }
      }
    }
  }
}
