resource "sdwan_eigrp_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  as_number    = 1
  address_families = [
    {
      type = "ipv4"
      redistributes = [
        {
          protocol     = "bgp"
          route_policy = "1.2.3.4"
        }
      ]
      networks = [
        {
          prefix = "1.2.3.4/24"
        }
      ]
    }
  ]
  hello_interval          = 5
  hold_time               = 15
  route_policy_name       = "RP1"
  filter                  = false
  authentication_type     = "hmac-sha-256"
  hmac_authentication_key = "myAuthKey"
  interfaces = [
    {
      interface_name = "Ethernet1"
      shutdown       = false
      summary_addresses = [
        {
          prefix = "1.2.3.4/24"
        }
      ]
    }
  ]
}
