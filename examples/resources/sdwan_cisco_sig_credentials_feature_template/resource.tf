resource "sdwan_cisco_sig_credentials_feature_template" "example" {
  name                     = "Example"
  description              = "My Example"
  device_types             = ["vedge-C8000V"]
  zscaler_organization     = "org1"
  zscaler_partner_base_uri = "abc"
  zscaler_username         = "user1"
  zscaler_password         = "password123"
  zscaler_cloud_name       = 1
  zscaler_partner_username = "partner1"
  zscaler_partner_password = "password123"
  zscaler_partner_api_key  = "key123"
  umbrella_api_key         = "key123"
  umbrella_api_secret      = "secret123"
  umbrella_organization_id = "org1"
}
