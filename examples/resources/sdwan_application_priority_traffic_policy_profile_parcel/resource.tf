resource "sdwan_application_priority_traffic_policy_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_action     = "accept"
  vpn                = ["1"]
  target_direction   = "all"
  sequences = [
    {
      sequence_id = 1
      name        = "RULE_1"
      base_action = "accept"
      protocol    = "ipv4"
      matches = [
        {
          dscp          = 1
          packet_length = "123"
          protocol      = ["2"]
          tcp           = "gre"
          traffic_to    = "core"
        }
      ]
      actions = [
        {
          counter = "COUNTER_1"
          log     = false
          sla_class = [
            {
              preferred_color       = ["default"]
              strict_drop           = true
              fallback_to_best_path = false
            }
          ]
          backup_sla_preferred_color = ["default"]
          sets = [
            {
              dscp                              = 1
              local_tloc_list_color             = ["default"]
              local_tloc_restrict               = "false"
              local_tloc_list_encapsulation     = "gre"
              tloc_ip                           = "1.2.3.4"
              tloc_color                        = ["default"]
              tloc_encapsulation                = "gre"
              service_type                      = "FW"
              service_color                     = ["default"]
              service_encapsulation             = "ipsec"
              service_tloc_ip                   = "1.2.3.4"
              service_vpn                       = "1"
              service_chain_type                = "SC1"
              service_chain_vpn                 = 1
              service_chain_local               = false
              service_chain_fallback_to_routing = false
              service_chain_tloc                = ["default"]
              service_chain_encapsulation       = "ipsec"
              service_chain_id                  = "1.2.3.4"
              next_hop                          = "1.2.3.4"
              next_hop_ipv6                     = "2001:0:0:1::/64"
              vpn                               = "1"
            }
          ]
          redirect_dns_field           = "redirectDns"
          redirect_dns_value           = "umbrella"
          tcp_optimization             = true
          dre_optimization             = true
          service_node_group           = "SNG-APPQOE1"
          loss_correction_type         = "fecAdaptive"
          loss_correct_fec_threshold   = 1
          cflowd                       = true
          nat_pool                     = 2
          nat_vpn                      = 0
          nat_fallback                 = false
          nat_bypass                   = false
          nat_dia_pool                 = [1]
          nat_dia_interface            = ["ethernet"]
          secure_internet_gateway      = true
          fallback_to_routing          = true
          secure_service_edge_instance = "zScaler"
        }
      ]
    }
  ]
}
