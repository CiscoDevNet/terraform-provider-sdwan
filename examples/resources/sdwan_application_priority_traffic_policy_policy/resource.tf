resource "sdwan_application_priority_traffic_policy_policy" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_action     = "accept"
  vpns               = ["Local_Internet_for_Guests"]
  direction          = "all"
  sequences = [
    {
      sequence_id   = 1
      sequence_name = "traffic"
      base_action   = "accept"
      protocol      = "ipv4"
      matches = [
        {
          dscp = 1
        }
      ]
      actions = [
        {
          set_parameters = [
            {
            }
          ]
        }
      ]
    }
  ]
}
