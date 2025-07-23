resource "sdwan_policy_group" "example" {
  name                = "PG_1"
  description         = "My policy group 1"
  solution            = "sdwan"
  feature_profile_ids = ["f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"]
}
