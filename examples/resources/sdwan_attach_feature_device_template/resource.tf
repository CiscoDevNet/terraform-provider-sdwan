resource "sdwan_attach_feature_device_template" "example" {
  id = sdwan_feature_device_template.DT1.id
  devices = [
    {
      id = "C8K-CC678D1C-8EDF-3966-4F51-ABFAB64F5ABE"
      variables = {
        system_site_id                                  = "1001"
        system_system_ip                                = "1.1.1.1"
        system_host_name                                = "router1"
        vpn_if_name_Default_vEdge_DHCP_Tunnel_Interface = "GigabitEthernet1"
      }
    }
  ]
}
