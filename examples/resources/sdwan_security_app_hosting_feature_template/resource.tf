resource "sdwan_security_app_hosting_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  virtual_application = [
    {
      instance_id        = "2e89c1fe-440a-43f5-9f3a-54a9836fdbb5"
      application_type   = "utd"
      nat                = true
      database_url       = false
      resource_profile   = "low"
      service_gateway_ip = "1.2.3.4/24"
      service_ip         = "1.2.3.4/24"
      data_gateway_ip    = "192.0.2.1/24"
      data_service_ip    = "192.0.2.1/24"
    }
  ]
}
