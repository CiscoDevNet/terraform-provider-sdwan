resource "sdwan_service_tracker_profile_parcel" "example" {
  name                  = "Example"
  description           = "My Example"
  feature_profile_id    = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  tracker_name          = "TRACKER_2"
  endpoint_api_url      = "google.com"
  endpoint_dns_name     = "google.com"
  endpoint_ip           = "1.2.3.4"
  protocol              = "tcp"
  port                  = 123
  interval              = 30
  multiplier            = 3
  threshold             = 300
  endpoint_tracker_type = "static-route"
  tracker_type          = "endpoint"
}
