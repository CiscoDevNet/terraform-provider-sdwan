resource "sdwan_service_multicast_feature" "example" {
  name                       = "Example"
  description                = "My Example"
  feature_profile_id         = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  spt_only                   = false
  local_replicator           = false
  local_replicator_threshold = 10
  igmp_interfaces = [
    {
      interface_name = "GigabitEthernet1"
      version        = 2
      join_groups = [
        {
          group_address  = "224.0.0.0"
          source_address = "1.2.3.4"
        }
      ]
    }
  ]
  pim_source_specific_multicast_enable      = true
  pim_source_specific_multicast_access_list = "1"
  pim_spt_threshold                         = "0"
  pim_interfaces = [
    {
      interface_name      = "GigabitEthernet1"
      query_interval      = 30
      join_prune_interval = 60
    }
  ]
  static_rp_addresses = [
    {
      ip_address  = "1.2.3.4"
      access_list = "1"
      override    = false
    }
  ]
  enable_auto_rp = false
  pim_bsr_rp_candidates = [
    {
      interface_name = "GigabitEthernet1"
      access_list_id = "2"
      interval       = 30
      priority       = 1
    }
  ]
  pim_bsr_candidates = [
    {
      interface_name               = "GigabitEthernet1"
      hash_mask_length             = 30
      priority                     = 120
      accept_candidate_access_list = "test"
    }
  ]
  msdp_groups = [
    {
      mesh_group_name = "Example"
      peers = [
        {
          peer_ip                      = "1.2.3.4"
          connection_source_interface  = "GigabitEthernet1"
          remote_as                    = 1
          peer_authentication_password = "Password123!"
          keepalive_interval           = 15
          keepalive_hold_time          = 30
          sa_limit                     = 1
          default_peer                 = false
        }
      ]
    }
  ]
  msdp_originator_id             = "GigabitEthernet1"
  msdp_connection_retry_interval = 30
}
