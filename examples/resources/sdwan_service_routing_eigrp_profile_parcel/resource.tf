resource "sdwan_service_routing_eigrp_profile_parcel" "example" {
  name                 = "Example"
  description          = "My Example"
  feature_profile_id   = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  autonomous_system_id = 111
  networks = [
    {
      ip_address = "100.2.2.3"
      mask       = "255.255.255.0"
    }
  ]
  hello_interval = 5
  hold_time      = 15
  type           = "md5"
  md5_keys = [
    {
      key_id             = 2
      authentication_key = "password123"
    }
  ]
  af_interfaces = [
    {
      name     = "GigabitEthernet3"
      shutdown = false
      summary_addresses = [
        {
          address = "10.0.0.1"
          mask    = "255.255.255.0"
        }
      ]
    }
  ]
  filter = false
}
