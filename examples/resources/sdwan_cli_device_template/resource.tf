resource "sdwan_cli_device_template" "example" {
  name              = "Example"
  description       = "My description"
  device_type       = "vedge-ISR-4331"
  cli_type          = "device"
  cli_configuration = " system\n host-name             R1-ISR4331-1200-1"
}
