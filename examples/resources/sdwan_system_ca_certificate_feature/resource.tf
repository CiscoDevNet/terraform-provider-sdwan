resource "sdwan_system_ca_certificate_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  certificates = [
    {
      trust_point_name  = "example"
      ca_certificate_id = "22a2f715-fed8-4031-b693-b5d43451a05e"
    }
  ]
}
