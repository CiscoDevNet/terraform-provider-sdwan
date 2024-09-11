resource "sdwan_transport_t1_e1_controller_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  type               = "t1"
  slot               = "11"
  entries = [
    {
      t1_description = "T1"
      t1_framing     = "esf"
      t1_linecode    = "ami"
      cable_length   = "long"
      length_long    = "-7.5db"
      clock_source   = "line"
      line_mode      = "primary"
      description    = "desc"
      channel_groups = [
        {
          channel_group = 12
          time_slot     = "timeslots 15"
        }
      ]
    }
  ]
}
