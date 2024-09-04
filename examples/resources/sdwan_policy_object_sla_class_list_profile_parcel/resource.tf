resource "sdwan_policy_object_sla_class_list_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  entries = [
    {
      latency                            = 2
      loss                               = 1
      jitter                             = 1
      fallback_best_tunnel_criteria      = "loss"
      fallback_best_tunnel_loss_variance = 5
    }
  ]
}
