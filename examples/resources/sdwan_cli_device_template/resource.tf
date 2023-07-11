resource "sdwan_cli_device_template" "test" {
  name              = "cli_template_1"
  description       = "Terraform integration test"
  device_type       = "vedge-ISR-4331"
  cli_type          = "device"
  cli_configuration = <<-EOT
    system
    host-name             {{test}}-ISR4331-1200-1
  EOT
}
