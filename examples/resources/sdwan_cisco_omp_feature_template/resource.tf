resource "sdwan_cisco_omp_feature_template" "example" {
  name                      = "Example"
  description               = "My Example"
  device_types              = ["vedge-C8000V"]
  graceful_restart          = true
  overlay_as                = 1
  send_path_limit           = 4
  ecmp_limit                = 4
  shutdown                  = false
  omp_admin_distance_ipv4   = 10
  omp_admin_distance_ipv6   = 10
  advertisement_interval    = 1
  graceful_restart_timer    = 43200
  eor_timer                 = 300
  holdtime                  = 60
  ignore_region_path_length = false
  transport_gateway         = "prefer"
  advertise_ipv4_routes = [
    {
      protocol                = "ospf"
      advertise_external_ospf = "external"
    }
  ]
  advertise_ipv6_routes = [
    {
      protocol = "ospf"
    }
  ]
}
