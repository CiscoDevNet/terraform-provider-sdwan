resource "sdwan_policy_object_app_probe_class" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  entries = [
    {
      map = [
        {
          color = "3g"
          dscp  = 45
        }
      ]
      forwarding_class = "classlist1"
    }
  ]
}
