resource "sdwan_cisco_bfd_feature_template" "example" {
  name          = "Example"
  description   = "My Example"
  device_types  = ["vedge-C8000V"]
  multiplier    = 3
  poll_interval = 800000
  default_dscp  = 48
  colors = [
    {
      color          = "private5"
      hello_interval = 1000
      multiplier     = 7
      pmtu_discovery = true
      dscp           = 46
    }
  ]
}
