resource "sdwan_transport_ipv6_acl_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_action     = "drop"
  sequences = [
    {
      sequence_id   = 1
      sequence_name = "AccessControlList1"
      conditions = [
        {
          next_header   = 10
          packet_length = 1500
          source_ports = [
            {
              port = 8000
            }
          ]
          tcp_state     = "syn"
          traffic_class = [10]
        }
      ]
      actions = [
        {
          accept_counter_name = "COUNTER_1"
          accept_log          = false
          accept_next_hop     = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
          traffic_class       = 10
        }
      ]
    }
  ]
}
