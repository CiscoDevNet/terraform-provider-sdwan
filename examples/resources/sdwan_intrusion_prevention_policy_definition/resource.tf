resource "sdwan_intrusion_prevention_policy_definition" "example" {
  name             = "Example"
  description      = "My description"
  mode             = "security"
  inspection_mode  = "protection"
  log_level        = "alert"
  custom_signature = false
  signature_set    = "connectivity"
  target_vpns      = ["1"]
  logging = [
    {
      external_syslog_server_ip  = "10.0.0.1"
      external_syslog_server_vpn = "123"
    }
  ]
}
