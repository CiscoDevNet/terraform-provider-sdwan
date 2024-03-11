resource "sdwan_eigrp_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  as_num       = 1
  address_family = [
    {
      type = "ipv4"
      redistribute = [
        {
          protocol     = "bgp"
          route_policy = "1.2.3.4"
        }
      ]
      network = [
        {
          prefix = "1.2.3.4/24"
        }
      ]
    }
  ]
  hello_interval          = 5
  hold_time               = 15
  route_policy_name       = "example"
  filter                  = false
  authentication_type     = "hmac-sha-256"
  hmac_authentication_key = "myAuthKey"
  interfaces = [
    {
      interface_name = "Ethernet"
      shutdown       = false
      summary_address = [
        {
          prefix = "1.2.3.4/24"
        }
      ]
    }
  ]
}
