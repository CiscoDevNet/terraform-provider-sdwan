resource "sdwan_topology_custom_control_feature" "example" {
  name                            = "Example"
  description                     = "My Example"
  feature_profile_id              = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_action                  = "reject"
  target_level                    = "SITE"
  target_inbound_hierarchy_uuids  = ["acb2ea53-4a95-4970-a1ab-9bac15edb961"]
  target_outbound_hierarchy_uuids = ["acb2ea53-4a95-4970-a1ab-9bac15edb961"]
  sequences = [
    {
      id          = 1
      name        = "Rule1"
      base_action = "accept"
      type        = "route"
      ip_type     = "ipv4"
      match_entries = [
        {
          omp_tag            = 100
          origin             = "connected"
          originator         = "1.2.3.4"
          hierarchy_uuids    = ["c446d770-2ac0-4e2c-9a64-345d562a4ac7"]
          tloc_ip            = "1.2.3.4"
          tloc_color         = "bronze"
          tloc_encapsulation = "ipsec"
        }
      ]
      action_entries = [
        {
          set_parameters = [
            {
              preference = 100
              omp_tag    = 100
            }
          ]
        }
      ]
    }
  ]
}
