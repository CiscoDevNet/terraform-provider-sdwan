resource "sdwan_cisco_thousandeyes_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  virtual_applications = [
    {
      instance_id            = "1"
      application_type       = "te"
      te_account_group_token = "1234567"
      te_vpn                 = 1
      te_agent_ip            = "1.1.1.2/24"
      te_default_gateway     = "1.1.1.255"
      te_name_server         = "10.2.2.2"
      te_hostname            = "agent1"
      te_web_proxy_type      = "static"
      te_proxy_host          = "3.3.3.3"
      te_proxy_port          = 80
    }
  ]
}
