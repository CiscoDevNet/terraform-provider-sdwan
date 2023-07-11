resource "sdwan_sla_class_policy_object" "example" {
  name = "Example"
  entries = [
    {
      jitter                        = 100
      latency                       = 10
      loss                          = 1
      fallback_best_tunnel_criteria = "jitter-loss-latency"
      fallback_best_tunnel_jitter   = 100
      fallback_best_tunnel_latency  = 10
      fallback_best_tunnel_loss     = 1
    }
  ]
}
