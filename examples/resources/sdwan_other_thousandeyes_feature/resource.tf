resource "sdwan_other_thousandeyes_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  virtual_application = [
    {
      account_group_token    = "qwer"
      vpn                    = 1
      management_ip          = "10.0.0.2"
      management_subnet_mask = "255.255.255.0"
      agent_default_gateway  = "10.0.0.1"
      name_server_ip         = "77.77.77.71"
      hostname               = "thousandeyesHost"
      proxy_type             = "static"
      proxy_host             = "proxy.thousandeyes.com"
      proxy_port             = 3128
    }
  ]
}
