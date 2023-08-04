resource "sdwan_intrusion_prevention_policy_definition" "example" {
  name            = "Example"
  description     = "My description"
  mode            = "security"
  inspection_mode = "protection"
  log_level       = "alert"
  signature_set   = "connectivity"
  target_vpns     = ["1"]
}
