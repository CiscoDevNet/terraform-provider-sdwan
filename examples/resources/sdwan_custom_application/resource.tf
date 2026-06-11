resource "sdwan_custom_application" "example" {
  app_name     = "Example"
  server_names = ["*customapp.com"]
  l3l4 = [
    {
      ip_addresses = ["192.168.1.0/24"]
      l4_protocol  = "TCP"
      ports        = "1 10-20"
    }
  ]
  application_family = "routing"
  application_group  = "ipsec-group"
  traffic_class      = "signaling"
  business_relevance = "business-relevant"
}
