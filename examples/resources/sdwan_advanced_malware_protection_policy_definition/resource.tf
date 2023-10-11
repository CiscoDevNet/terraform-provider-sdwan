resource "sdwan_advanced_malware_protection_policy_definition" "example" {
  name                        = "Example"
  description                 = "My description"
  mode                        = "security"
  match_all_vpn               = false
  target_vpns                 = ["1"]
  alert_log_level             = "critical"
  amp_cloud_region            = "apjc"
  amp_cloud_region_est_server = "apjc"
  file_analysis               = false
}
