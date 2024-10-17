resource "sdwan_transport_ipv4_acl_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_action     = "drop"
  sequences = [
    {
      sequence_id   = 1
      sequence_name = "AccessControlList1"
      match_entries = [
        {
          dscps         = [16]
          packet_length = 1500
          protocols     = [1]
          source_ports = [
            {
              port = 8000
            }
          ]
          tcp_state = "syn"
        }
      ]
      actions = [
        {
          accept_set_dscp     = 60
          accept_counter_name = "COUNTER_1"
          accept_log          = false
          accept_set_next_hop = "1.2.3.4"
        }
      ]
    }
  ]
}
